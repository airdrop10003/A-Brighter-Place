// A Brighter Place - Typescript

// A function to generate a party banner
function makePartyBanner(name: string) {
  let banner = "";
  // Create the banner using a for loop
  for (let i = 0; i < name.length; i++) {
    banner += "*";
  }
  // Add the name to the banner
  banner += `\n* ${name} *\n`;
  // Finish the banner
  for (let i = 0; i < name.length; i++) {
    banner += "*";
  }
  return banner;
}

// A function to generate a warm welcome
function giveWelcome(name: string) {
  return `Welcome, ${name}! It's so good to have you here.`;
}

// A function to generate a farewell message
function sendOff(name: string) {
  return `Thank you, ${name}! Please come again soon.`;
}

// A function to calculate the area of a circle
function areaOfCircle(radius: number) {
  const pi = 3.14159;
  return pi * Math.pow(radius, 2);
}

// A function to count the number of vowels in a string
function countVowels(word: string) {
  let count = 0;
  // Use a for loop to check each character
  for (let i = 0; i < word.length; i++) {
    // Catch vowels
    if (word[i] === "a" || word[i] === "e" || word[i] === "i" || word[i] === "o" || word[i] === "u") {
      count++;
    }
  }
  return count;
}

// A function to rotate a two-dimensional array 90 degrees
function rotate90(matrix: number[][]) {
  const rows = matrix.length;
  const cols = matrix[0].length;
  const result = [];
  // Iterate the columns
  for (let j = 0; j < cols; j++) {
    const row = [];
    // Iterate the rows
    for (let i = rows - 1; i >= 0; i--) {
      row.push(matrix[i][j]);
    }
    result.push(row);
  }
  return result;
}

// A function to reverse a string
function reverseString(value: string) {
  let reverse = "";
  // Use a for loop to reverse the string
  for (let i = value.length - 1; i >= 0; i--) {
    reverse += value[i];
  }
  return reverse;
}

// A function to calculate the average of an array of numbers
function arrayAverage(numbers: number[]) {
  let total = 0;
  // Use a for loop to add each number to the total
  for (let i = 0; i < numbers.length; i++) {
    total += numbers[i];
  }
  // Divide the total by the length of the array
  return total / numbers.length;
}

// A function to calculate the factorial of a number
function factorial(num: number) {
  let total = 1;
  // Use a for loop to calculate the factorial
  for (let i = num; i > 0; i--) {
    total *= i;
  }
  return total;
}

// A function to calculate the Fibonacci sequence
function fibonacci(num: number) {
  let a = 0;
  let b = 1;
  let result = [a, b];
  // Use a for loop to calculate the next number in the sequence
  for (let i = 2; i <= num; i++) {
    let c = a + b;
    a = b;
    b = c;
    result.push(c);
  }
  return result;
}

// A function to check if a number is prime
function isPrime(num: number) {
  for (let i = 2; i < num; i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return num > 1;
}

// A function to calculate the maximum value in an array of numbers
function max(numbers: number[]) {
  let maxValue = numbers[0];
  // Use a for loop to find the maximum value
  for (let i = 1; i < numbers.length; i++) {
    if (numbers[i] > maxValue) {
      maxValue = numbers[i];
    }
  }
  return maxValue;
}

// A function to capitalize the first letter of a string
function capitalize(word: string) {
  return word[0].toUpperCase() + word.slice(1);
}

// A function to remove duplicate values from an array
function uniqueValues(numbers: number[]) {
  let unique = [];
  // Use a for loop to identify unique numbers
  for (let i = 0; i < numbers.length; i++) {
    if (!unique.includes(numbers[i])) {
      unique.push(numbers[i]);
    }
  }
  return unique;
}

// A function to sort an array of strings in alphabetical order
function alphabeticalOrder(words: string[]) {
  return words.sort();
}

// A function to check if two strings are anagrams
function areAnagrams(str1: string, str2: string) {
  // Check that the strings are the same length
  if (str1.length !== str2.length) {
    return false;
  }
  // Concatenate the strings and sort them
  const sorted1 = str1.split("").sort().join("");
  const sorted2 = str2.split("").sort().join("");
  // Check if the sorted strings are the same
  return sorted1 === sorted2;
}

// A function to convert a string to an array of characters
function stringToArray(value: string) {
  return value.split("");
}

// A function to convert a number to a binary string
function numberToBinary(num: number) {
  let result = "";
  // Use a while loop to find the binary representation
  while (num > 0) {
    result = (num % 2) + result;
    num = Math.floor(num / 2);
  }
  return result;
}

// A function to find the sum of two numbers
function sum(a: number, b: number) {
  return a + b;
}

// A function to check if two strings are palindromes
function isPalindrome(str1: string, str2: string) {
  // Reverse one of the strings
  const reverse = str2
    .split("")
    .reverse()
    .join("");
  // Check if the reversed string is the same
  return str1 === reverse;
}

// A function to convert a string to lowercase
function toLowerCase(word: string) {
  return word.toLowerCase();
}

// A function to count the number of occurrences of a character in a string
function countOccurrences(word: string, char: string) {
  let count = 0;
  // Use a for loop to count the number of occurrences
  for (let i = 0; i < word.length; i++) {
    if (word[i] === char) {
      count++;
    }
  }
  return count;
}

// A function to count the number of words in a string
function countWords(sentence: string) {
  // Split the string into an array of words
  const words = sentence.split(" ");
  return words.length;
}

// A function to calculate the power of a number
function power(base: number, exponent: number) {
  return Math.pow(base, exponent);
}

// A function to calculate the distance between two points
function calculateDistance(x1: number, y1: number, x2: number, y2: number) {
  const a = x2 - x1;
  const b = y2 - y1;
  return Math.sqrt(a * a + b * b);
}

// A function to calculate the area of a triangle
function triangleArea(a: number, b: number, c: number) {
  const s = (a + b + c) / 2;
  return Math.sqrt(s * (s - a) * (s - b) * (s - c));
}

// A function to determine if a number is even
function isEven(num: number) {
  return num % 2 === 0;
}

// A function to generate a random number between two values
function randomNumber(min: number, max: number) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

// A function to generate a random hexadecimal color code
function randomColor() {
  const num1 = randomNumber(0, 255).toString(16);
  const num2 = randomNumber(0, 255).toString(16);
  const num3 = randomNumber(0, 255).toString(16);
  return `#${num1}${num2}${num3}`;
}

// A function to find the maximum depth of a nested object
function maxDepth(obj: any) {
  let level = 1;
  for (const key in obj) {
    if (typeof obj[key] == "object") {
      const subLevel = maxDepth(obj[key]) + 1;
      level = Math.max(subLevel, level);
    }
  }
  return level;
}

// A function to find the sum of digits of a number
function sumDigits(num: number) {
  let total = 0;
  // Use a while loop to add the digits of the number
  while (num > 0) {
    total += num % 10;
    num = Math.floor(num / 10);
  }
  return total;
}

// A function to check if a string is a valid email address
function isValidEmail(email: string) {
  const regex = /\S+@\S+\.\S+/;
  return regex.test(email);
}

// A function to calculate the sum of items in an array
function arraySum(numbers: number[]) {
  let total = 0;
  // Use a for loop to add the items in the array
  for (let i = 0; i < numbers.length; i++) {
    total += numbers[i];
  }
  return total;
}

// A function to calculate the most common element in an array
function mostCommonElement(numbers: number[]) {
  // Create an object to count the occurrences of each number
  const counts = {};
  numbers.forEach(num => {
    if (counts[num]) {
      counts[num]++;
    } else {
      counts[num] = 1;
    }
  });
  // Find the most common element
  let mostCommon = 0;
  let mostCommonCount = 0;
  for (const num in counts) {
    if (counts[num] > mostCommonCount) {
      mostCommon = Number(num);
      mostCommonCount = counts[num];
    }
  }
  return mostCommon;
}

// A function to check if an array is sorted
function isSorted(numbers: number[]) {
  let currentNum = numbers[0];
  // Use a for loop to check if each number is greater than the previous
  for (let i = 1; i < numbers.length; i++) {
    if (numbers[i] < currentNum) {
      return false;
    }
    currentNum = numbers[i];
  }
  return true;
}

// A function to find the second largest number in an array
function secondLargest(numbers: number[]) {
  let largest = numbers[0];
  let second = 0;
  // Use a for loop to find the second largest number
  for (let i = 1; i < numbers.length; i++) {
    if (numbers[i] > largest) {
      second = largest;
      largest = numbers[i];
    } else if (numbers[i] > second && numbers[i] !== largest) {
      second = numbers[i];
    }
  }
  return second;
}

// A function to calculate the Levenshtein distance between two strings
function levenshteinDistance(str1: string, str2: string) {
  // Create a matrix to store the values
  const matrix = [];
  // Initialize the first row
  for (let i = 0; i <= str2.length; i++) {
    matrix[i] = [i];
  }
  // Initialize the first column
  for (let j = 0; j <= str1.length; j++) {
    matrix[0][j] = j;
  }
  // Fill in the rest of the matrix
  for (let i = 1; i <= str2.length; i++) {
    for (let j = 1; j <= str1.length; j++) {
      if (str2[i - 1] === str1[j - 1]) {
        matrix[i][j] = matrix[i - 1][j - 1];
      } else {
        matrix[i][j] = Math.min(matrix[i - 1][j - 1] + 1, matrix[i][j - 1] + 1, matrix[i - 1][j] + 1);
      }
    }
  }
  // Return the last value in the matrix
  return matrix[str2.length][str1.length];
}

// A function to find the most frequent item in an array
function mostFrequent(numbers: number[]) {
  // Create an object to count the occurrences of each number
  const counts = {};
  numbers.forEach(num => {
    if (counts[num]) {
      counts[num]++;
    } else {
      counts[num] = 1;
    }
  });
  // Find the most frequent item
  let mostFrequentItem = 0;
  let mostFrequentCount = 0;
  for (const num in counts) {
    if (counts[num] > mostFrequentCount) {
      mostFrequentItem = Number(num);
      mostFrequentCount = counts[num];
    }
  }
  return mostFrequentItem;
}

// A function to check if a number is a perfect square
function isPerfectSquare(num: number) {
  return Math.sqrt(num) % 1 ===