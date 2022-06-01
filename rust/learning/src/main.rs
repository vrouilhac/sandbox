use std::io;

fn main() {
    let chars = ('│', '─', '╭', '╮', '╯', '╰');

    println!("Enter your name :");

    let mut username = String::new();
    io::stdin().read_line(&mut username).expect("Failed to enter username");
    let mut hello = String::from("Hello ");
    hello.push_str(&username);

    let top_row = build_row(&hello, (chars.2, chars.1, chars.3), true);
    let top_row: String = top_row.into_iter().collect();

    let mid_row = build_row(&hello, (chars.0, chars.1, chars.0), false);
    let mid_row: String = mid_row.into_iter().collect();

    let bot_row = build_row(&hello, (chars.5, chars.1, chars.4), true);
    let bot_row: String = bot_row.into_iter().collect();

    println!("{}", top_row);
    println!("{}", mid_row);
    println!("{}", bot_row);

}

fn build_row(string: &str, chars: (char, char, char), is_border: bool) -> Vec<char> {
    let mut row: Vec<char> = Vec::new();
    let length = string.len() + 1;
    let mut col = 0;

    loop {
        if col == 0 {
            row.push(chars.0);
        } else if col == length - 1 {
            row.push(chars.2);
        } else if is_border {
            row.push(chars.1);
        } else {
            let str_char = string.chars().nth(col-1).unwrap();
            row.push(str_char);
        }

        col = col + 1;

        if col == length {
            break;
        }
    };

    row
}
