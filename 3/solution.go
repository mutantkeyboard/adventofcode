package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func multiply(a, b int) int {
    return a * b
}

type instruction struct {
    typ     string // "mul", "do", "dont"
    pos     int
    num1    int
    num2    int
    rawText string
}

func extractMulExpressions(scanner *bufio.Scanner) int {
    mulPattern := `mul\(([0-9]+),([0-9]+)\)`
    doPattern := `do\(\)`
    dontPattern := `don't\(\)`
    
    mulRegex := regexp.MustCompile(mulPattern)
    doRegex := regexp.MustCompile(doPattern)
    dontRegex := regexp.MustCompile(dontPattern)
    
    sum := 0
    enabled := true // mul operations start enabled
    lineNum := 0

    for scanner.Scan() {
        line := scanner.Text()
        lineNum++
        
        // Find all instructions in the line
        var instructions []instruction
        
        // Find mul instructions
        mulMatches := mulRegex.FindAllStringSubmatchIndex(line, -1)
        for _, match := range mulMatches {
            num1, _ := strconv.Atoi(line[match[2]:match[3]])
            num2, _ := strconv.Atoi(line[match[4]:match[5]])
            instructions = append(instructions, instruction{
                typ:     "mul",
                pos:     match[0],
                num1:    num1,
                num2:    num2,
                rawText: line[match[0]:match[1]],
            })
        }
        
        // Find do() instructions
        doMatches := doRegex.FindAllStringIndex(line, -1)
        for _, match := range doMatches {
            instructions = append(instructions, instruction{
                typ:     "do",
                pos:     match[0],
                rawText: line[match[0]:match[1]],
            })
        }
        
        // Find don't() instructions
        dontMatches := dontRegex.FindAllStringIndex(line, -1)
        for _, match := range dontMatches {
            instructions = append(instructions, instruction{
                typ:     "dont",
                pos:     match[0],
                rawText: line[match[0]:match[1]],
            })
        }

        // Sort instructions by position
        // Using a simple bubble sort since we expect small number of instructions per line
        for i := 0; i < len(instructions)-1; i++ {
            for j := 0; j < len(instructions)-i-1; j++ {
                if instructions[j].pos > instructions[j+1].pos {
                    instructions[j], instructions[j+1] = instructions[j+1], instructions[j]
                }
            }
        }
        
        // Process instructions in order
        for _, inst := range instructions {
            switch inst.typ {
            case "do":
                enabled = true
                fmt.Printf("Line %d: Enabled at position %d (%s)\n", lineNum, inst.pos, inst.rawText)
            case "dont":
                enabled = false
                fmt.Printf("Line %d: Disabled at position %d (%s)\n", lineNum, inst.pos, inst.rawText)
            case "mul":
                if enabled {
                    result := multiply(inst.num1, inst.num2)
                    sum += result
                    fmt.Printf("Line %d: Added multiplication at pos %d: %d * %d = %d (sum: %d)\n", 
                        lineNum, inst.pos, inst.num1, inst.num2, result, sum)
                } else {
                    fmt.Printf("Line %d: Skipped disabled multiplication at pos %d: %s\n", 
                        lineNum, inst.pos, inst.rawText)
                }
            }
        }
        
        // Print line state summary
        fmt.Printf("Line %d state: enabled=%v, instructions=%d, lineSum=%d\n", 
            lineNum, enabled, len(instructions), sum)
        fmt.Println(strings.Repeat("-", 80))
    }

    return sum
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    total := extractMulExpressions(scanner)

    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Final sum of enabled multiplications: %d\n", total)
}