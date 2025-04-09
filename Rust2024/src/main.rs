use std::fs::read_to_string;

fn main(){
    let mut left = Vec::new();
    let mut right = Vec::new();

   for line in read_to_string("src/Input.txt").unwrap().lines() {
        if line.trim().is_empty() {
            continue;
        }
        let numbers: Vec<&str> = line.split_whitespace().collect();
        left.push(numbers[0].parse::<i64>().unwrap());
        right.push(numbers[1].parse::<i64>().unwrap());
    }

    left.sort();
    right.sort();

    let mut difference = Vec::new();
    for i in 0..left.len() {
        difference.push(left[i].abs()- right[i].abs());
    }

    let difference_sum: i64 = difference.iter().sum();
    print!("result: {:?}", difference_sum);
}
