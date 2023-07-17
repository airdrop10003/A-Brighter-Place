import scala.collection.mutable.ListBuffer
import java.util.Calendar

object ABrighterPlace {
  def main(args: Array[String]): Unit = {
    // Defines the base set of colours used in the program
    val primaryColors = List("Red", "Green", "Blue")
    
    // Constructs the ListBuffer to hold the mixedColors
    val mixedColors = ListBuffer[String]()
    
    // Iterates over the primaryColors to create the mixedColors
    for (primaryColor <- primaryColors) {
      val secondaryColors = primaryColor match {
        case "Red" => List("Pink", "Crimson")
        case "Green" => List("Lime", "Jade")
        case "Blue" => List("Cyan", "Azure")
      }
      for (secondaryColor <- secondaryColors) {
        mixedColors += primaryColor + "-" + secondaryColor
      }
    }
    
    // Prints out the mixedColors
    println(mixedColors)
    
    // Creates a Calendar instance
    val cal = Calendar.getInstance()
    
    // Gets the current hour of the day
    val currentHour = cal.get(Calendar.HOUR_OF_DAY)
    
    // Customizes the greeting according to the time of day
    val greeting = currentHour match {
      case x if x < 11 => "Good morning!"
      case x if x < 17 => "Good afternoon!"
      case _ => "Good evening!"
    }
    
    // Prints out the greeting
    println(greeting)
    
    // Defines the functions for calculating the area and circumference
    def area(radius: Double) = math.Pi * math.pow(radius, 2)
    def circumference(radius: Double) = 2 * math.Pi * radius
    
    // Calculates the area and circumference of a circle with radius 5
    val areaResult = area(5)
    val circumferenceResult = circumference(5)
    
    // Prints out the results
    println(s"The area of a circle with radius 5 is $areaResult")
    println(s"The circumference of a circle with radius 5 is $circumferenceResult")
    
    // Constructs a list of numbers
    val numbers = List(1, 2, 3, 4, 5)
    
    // Computes the sum of the numbers
    val sumResult = numbers.sum
    
    // Prints out the result
    println(s"The sum of the numbers is $sumResult")
    
    // Constructs a list of strings
    val words = List("Lorem", "ipsum", "dolor", "sit", "amet")
    
    // Computes the length of the longest string
    val longestLength = words.map(_.length).max
    
    // Prints out the result
    println(s"The length of the longest word is $longestLength")
    
    // Constructs a list of tuples
    val coordinates = List((1,2), (3,4), (5,6))
    
    // Computes the sum of the x coordinates
    val xCoordinateSum = coordinates.map(_._1).sum
    
    // Prints out the result
    println(s"The sum of the x coordinates is $xCoordinateSum")
    
    // Loops through the range 0 to 10
    for (i <- 0 to 10) {
      // Prints out the result
      println(s"The value of i is $i")
    }
    
    // Constructs a list of names
    val names = List("John", "Jane", "Peter", "Paul")
    
    // Filters out the names with length greater than 4
    val longNames = names.filter(_.length > 4)
    
    // Prints out the result
    println(s"The long names are $longNames")
    
    // Constructs a Map
    val ageMap = Map("John" -> 25, "Jane" -> 23, "Peter" -> 30)
    
    // Gets the age of John
    val johnsAge = ageMap("John")
    
    // Prints out the result
    println(s"John's age is $johnsAge")
    
    // Constructs a list of tuples
    val numbersPairs = List((2,3), (4,5), (6,7))
    
    // Computes the sum of the numbers in each tuple
    val sums = numbersPairs.map(x => x._1 + x._2)
    
    // Prints out the result
    println(s"The sums are $sums")
    
    // Defines the function for calculating the factorial
    def factorial(n: Int): Int = {
      if (n == 0) 1 else n * factorial(n - 1)
    }
    
    // Calculates the factorial of 5
    val result = factorial(5)
    
    // Prints out the result
    println(s"The factorial of 5 is $result")
    
    // Constructs a List of numbers
    val numbersList = List(1, 2, 3, 4, 5)
    
    // Computes the product of the numbers in the list
    val productResult = numbersList.product
    
    // Prints out the result
    println(s"The product of the numbers is $productResult")
    
    // Constructs a string
    val myString = "Hello World!"
    
    // Computes the length of the string
    val lengthResult = myString.length
    
    // Prints out the result
    println(s"The length of the string is $lengthResult")
    
    // Defines the function for calculating the sum of two numbers
    def sum(x: Int, y: Int): Int = x + y
    
    // Calculates the sum of 3 and 4
    val sumResult = sum(3, 4)
    
    // Prints out the result
    println(s"The sum of 3 and 4 is $sumResult")
    
    // Constructs a list of names
    val namesList = List("John", "Jane", "Peter", "Paul")
    
    // Filters out the names with length less than 5
    val shortNames = namesList.filter(_.length < 5)
    
    // Prints out the result
    println(s"The short names are $shortNames")
    
    // Constructs a sequence of numbers
    val nums = Seq(1, 2, 3, 4, 5)
    
    // Computes the average of the numbers in the sequence
    val averageResult = nums.sum.toDouble / nums.length
    
    // Prints out the result
    println(s"The average of the numbers is $averageResult")
    
    // Constructs a Map
    val nameMap = Map("John" -> 25, "Jane" -> 23, "Peter" -> 30)
    
    // Computes the average age
    val averageAge = nameMap.values.sum.toDouble / nameMap.size
    
    // Prints out the result
    println(s"The average age is $averageAge")
    
    // Constructs a tuple
    val firstTuple = (1, "John")
    
    // Gets the first element of the tuple
    val firstElement = firstTuple._1
    
    // Prints out the result
    println(s"The first element of the tuple is $firstElement")
    
    // Constructs a Set
    val numSet = Set(1, 2, 3, 4, 5)
    
    // Computes the union of the Set and a List
    val resultSet = numSet ++ List(6, 7, 8)
    
    // Prints out the result
    println(s"The union of the Set and the List is $resultSet")
    
    // Constructs a sequence of words
    val wordsSeq = Seq("Lorem", "ipsum", "dolor", "sit", "amet")
    
    // Returns a new sequence with the words in reversed order
    val reversedWords = wordsSeq.reverse
    
    // Prints out the result
    println(s"The reversed words are $reversedWords")
    
    // Constructs a List
    val numList = List(1, 2, 3, 4, 5)
    
    // Computes the sum of the numbers multiplied by two
    val sumByTwo = numList.map(_ * 2).sum
    
    // Prints out the result
    println(s"The sum of the multiplied numbers is $sumByTwo")
    
    // Constructs a Map
    val ageMap2 = Map("John" -> 25, "Jane" -> 23, "Peter" -> 30)
    
    // Gets the names of people older than 25
    val olderThan25 = ageMap2.filter(_._2 > 25).keys
    
    // Prints out the result
    println(s"The names of people older than 25 are $olderThan25")
    
    // Constructs a list of tuples
    val numbersPairs2 = List((2,3), (4,5), (6,7))
    
    // Computes the difference of the numbers in each tuple
    val differences = numbersPairs2.map(x => x._1 - x._2)
    
    // Prints out the result
    println(s"The differences are $differences")
    
    // Creates a Range
    val myRange = 0 to 10
    
    // Computes the sum of the numbers in the range
    val rangeSum = myRange.sum
    
    // Prints out the result
    println(s"The sum of the numbers in the range is $rangeSum")
    
    // Constructs a list of numbers
    val numbersList2 = List(1, 2, 3, 4, 5)
    
    // Returns a list with the numbers multiplied by two
    val doubled = numbersList2.map(_ * 2)
    
    // Prints out the result
    println(s"The doubled numbers are $doubled")
    
    // Constructs a string
    val myString2 = "Hello World!"
    
    // Returns a new string with all the characters in reversed order
    val reversedString = myString2.reverse
    
    // Prints out the result
    println(s"The reversed string is $reversedString")
    
    // Constructs a Set
    val numSet2 = Set(1, 2, 3, 4, 5)
    
    // Computes the intersection of the Set and a List
    val resultSet2 = numSet2.intersect(List(4, 5, 6, 7))
    
    // Prints out the result
    println(s"The intersection of the Set and the List is $resultSet2")
  }
}