// Imports
use std::io;

fn main() {
  // Prints out 'A Brighter Place'
  println!("A Brighter Place");

  // Prompt user to enter name
  println!("Please enter your name: ");

  // Read user input
  let mut name = String::new();
  io::stdin().read_line(&mut name)
    .expect("Failed to read line");

  // Prints out name
  println!("Hello {}", name);

  // Create two integer variables
  let num1 = 10;
  let num2 = 5;

  // Perform addition
  let result = num1 + num2;
  println!("The sum of {} and {} is: {}", num1 ,num2, result);

  // Create three booleans
  let is_true = true;
  let is_false = false;
  let is_allowed = true;

  // Perform logical AND
  let result = is_true && is_allowed;
  println!("The result of AND operation is: {}", result);

  // Create an array
  let array = [1, 2, 3, 4, 5];

  // Iterate through array
  for element in array.iter() {
    // Print all elements in the array
    println!("Array element: {}", element);
  }

  // Create a loop
  for i in 0..5 {
    // Prints numbers from 0 to 4
    println!("Number: {}", i);
  }

  // Create a new function
  fn divide(a:i32, b:i32) -> i32 {
   let result = a / b;
   result
  }

  // Call function to perform division
  let division_result = divide(20, 4);
  println!("The result of division is: {}", division_result);

  // Use match expression
  let number = 5;
  match number {
    1 => println!("The number is 1"),
    2 => println!("The number is 2"),
    3 => println!("The number is 3"),
    4 => println!("The number is 4"),
    5 => println!("The number is 5"),
    _ => println!("The number is not 1-5"),
  }

  // Create a string
  let text = String::from("We are on a journey to a brighter place");

  // Iterate through each character
  for character in text.chars() {
    // Print out each character
    println!("{}", character);
  }

  // Create a vector
  let mut vector = vec![10, 20, 30, 40];

  // Push an element to the vector
  vector.push(50);

  // Print out the vector
  println!("{:?}", vector);

  // Create a hash map
  use std::collections::HashMap;
  let mut map = HashMap::new();

  // Insert elements to hash map
  map.insert("A", 1);
  map.insert("B", 2);
  map.insert("C", 3);

  // Iterate through map
  for (key, value) in &map {
    // Print out each key and value
    println!("{}: {}", key, value);
  }

  // Create a struct
  struct Point {
    x: i32,
    y: i32,
  }

  // Create an instance of the struct
  let point = Point {x: 10, y: 20};

  // Print out the instance's fields
  println!("Point: ({}, {})", point.x, point.y);

  // Create an enum
  enum Color {
    Red,
    Green,
    Blue,
  }

  // Create an instance of the enum
  let color = Color::Blue;

  // Match the enum instance
  match color {
    Color::Red => println!("The color is Red"),
    Color::Green => println!("The color is Green"),
    Color::Blue => println!("The color is Blue"),
  }

  // Create a trait
  trait Animal {
    fn sound(&self) -> String;
  }

  // Create a struct that implements the trait
  struct Dog {
    name: String,
  }

  // Implement the trait for the struct
  impl Animal for Dog {
    fn sound(&self) -> String {
      // Return the bark sound
      String::from("Woof")
    }
  }

  // Create an instance of the dog
  let dog = Dog {name: String::from("Max") };

  // Print out the dog's sound
  println!("{} says {}", dog.name, dog.sound());

  // Create a closure
  let add_nums = |x: i32, y: i32| x + y;

  // Call closure to perform addition
  let result = add_nums(10, 20);
  println!("The result of addition is: {}", result);

  // Create a reference
  let a = 10;
  let b = &a;

  // Print out the reference
  println!("b= {}", b);

  // Check for power-of-two
  fn is_power_of_two(num: i32) -> bool {
    if num == 0 {
      return false;
    }
    
    let mut result = num;
    while result != 1 {
      if result % 2 != 0 {
        return false;
      }
      result /= 2;
    }
    true
  }

  // Check if 16 is a power of two
  let result = is_power_of_two(16);
  println!("Is 16 a power of two: {}", result);

  // Perform bitwise operations
  let a = 10;
  let b = 5;

  // Perform bitwise AND
  let result = a & b;
  println!("The result of AND operation is: {}", result);

  // Perform bitwise OR
  let result = a | b;
  println!("The result of OR operation is: {}", result);

  // Perform bitwise XOR
  let result = a ^ b;
  println!("The result of XOR operation is: {}", result);

  // Create a generic function
  fn get_elements<T>(arr: [T; 5]) -> Vec<T> {
    let mut result = Vec::new();
    
    for element in arr.iter() {
      result.push(*element);
    }

    result
  }

  // Create an array
  let arr = [1, 2, 3, 4, 5];

  // Call the generic function
  let result = get_elements(arr);
  println!("The result of the generic function is: {:?}", result);

  // Use the ? operator
  fn divide_with_option(a: i32, b: i32) -> Option<i32> {
    if b == 0 {
      None
    } else {
      Some(a / b)
    }
  }

  // Call the function with ?
  let result = divide_with_option(20, 4)?;
  println!("The result of division is: {}", result);

  // Unpack a tuple
  let tuple = ("Max", 10);
  let (name, age) = tuple;
  println!("name={}, age={}", name, age);

  // Use a type alias
  type Name = String;
  let name: Name = String::from("John");
  println!("Name={}", name);

  // Create a macro
  macro_rules! say_hello {
    () => (
      println!("Hello!")
    )
  }

  // Call the macro
  say_hello!();
}