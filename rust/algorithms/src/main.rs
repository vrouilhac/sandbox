use std::io;

fn compute_fibonacci(limit: i32) -> i32 {
    let mut counter = 0;
    let mut minus1 = 0;
    let mut minus2 = 0;
    let mut result = 0;

    loop {
        if counter >= limit {
            break;
        }

        if result == 0 {
            minus1 = 0;
            minus2 = 1;
        }

        minus1 = minus2.clone();

        counter += 1;
    }

    result
}

fn main() {
    let mut limit = String::new();    
    
    println!("Choose a number : ");
    io::stdin().read_line(&mut limit).expect("Something wrong reading input");

    let limit: i32 = match limit.trim().parse::<i32>() {
        Ok(limit) => limit,
        Err(_) => 0,
    };
    let fibonacci_number = compute_fibonacci(limit.clone());

    println!("The fibonacci number for {} is {}", limit, fibonacci_number);
}
