use std::collections::HashMap;

fn main() {
  let mut map = HashMap::new();
  let vincent_name = String::from("vincent");

  map.insert(&vincent_name, 43);

  println!("The value for {} is {}", &vincent_name, match map.get(&vincent_name) {
    Some(value) => *value,
    None => -1,
  });
}

// Need to learn lifetime dude
// fn comput() -> HashMap<&str, i32> {
//     let mut testing = HashMap::new();
//     testing.insert("test", 23);
//     testing
// }
