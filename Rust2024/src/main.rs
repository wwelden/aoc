use std::fs;
use std::io;



fn readFile() -> io::Result<()> {
    let file_path = "src/input.txt";
    let contents = fs::read_to_string(file_path)?;
    println!("File contents:\n{}", contents);
    Ok(())
}

fn parseInput(input: string)-> Vec<[i32; 2]>> {
    let lines: Vec<&str> = inout.trim().lines().collect();
    let mut result = vec![[0; 2]; lines.len()];

    for (i, line) in line.iter().enumerate() {
        let numbers: Vec<&str> = line.split_whitespace().collect();
        if numbers.len() == 2 {
            let num1 = numbers[0].parse::<i32>().unwrap_or(0);
            let num2 = numbers[1].parse::<i32>().unwrap_or(0);
            result[i] = [num1, num2];
        }
    }

    result
}
