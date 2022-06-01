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
