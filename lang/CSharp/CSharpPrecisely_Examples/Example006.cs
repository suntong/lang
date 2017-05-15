// Example 6 from page 5 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

class LayoutExample {                   // Class declaration
  int j;

  LayoutExample(int j) {
    this.j = j;                         // One-line body
  }

  int Sum(int b) {                      // Multi-line body
    if (j > 0) {                        // If statement
      return j + b;                     // Single statement
    } else if (j < 0) {                 // Nested if-else, block statement
      int res = -j + b;
      return res * 117;
    } else { // j == 0                  // Terminal else, block statement
      int sum = 0;
      for (int i=0; i<10; i++)          // For loop
        sum += (b - i) * (b - i);
      return sum;
    }
  }

  static void Main() { 
  }
}
