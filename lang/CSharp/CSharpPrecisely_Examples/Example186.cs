// Example 186 from page 151 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.IO;                // StringReader, TextReader
using System.Text;              // StringBuilder

class MyTest {
  public static void Main(String[] args) {
    if (args.Length == 1) 
      Tokenize(new StringReader(args[0]));
    else
      Tokenize(new StringReader("(6 + abc2) * 3343"));
  }
	
  public static void Tokenize(TextReader rd) {
    while (rd.Peek() != -1) {
      if (Char.IsWhiteSpace((char)rd.Peek()))           // Whitespace, skip
        rd.Read();
      else if (Char.IsDigit((char)rd.Peek())) {         // Number
        int val = rd.Read() - '0';
        while (Char.IsDigit((char)rd.Peek())) 
          val = 10 * val + rd.Read() - '0';
        Console.WriteLine(new Int(val));
      } else if (Char.IsLetter((char)rd.Peek())) {      // Identifier
        StringBuilder id = new StringBuilder().Append((char)rd.Read());
        while (Char.IsLetterOrDigit((char)rd.Peek())) 
          id.Append((char)rd.Read());
        Console.WriteLine(new Id(id.ToString()));
      } else 
        switch (rd.Peek()) {
        case '+': case '-': case '*': case '/':         // Operator
          Console.WriteLine(new Op((char)rd.Read())); break;
        case '(': case ')':                             // Separator
          Console.WriteLine(new Sep((char)rd.Read())); break;
        default:                                        // Illegal token
          throw new ApplicationException("Illegal character '"+(char)rd.Peek()+"'");
        }
    }
  }
}

// Classes to represent tokens: identifiers, numbers, operators, delimiters

abstract class Token { }

class Int : Token { 
  public readonly int i;
  
  public Int(int i) {
    this.i = i;
  }

  public override String ToString() {
    return String.Format("int:{0}", i);
  }
}

class Id : Token { 
  public readonly String id;
  
  public Id(String id) {
    this.id = id;
  }

  public override String ToString() {
    return String.Format("id:{0}", id);
  }
}

class Op : Token { 
  public readonly char op;
  
  public Op(char op) {
    this.op = op;
  }

  public override String ToString() {
    return String.Format("op:{0}", op);
  }
}

class Sep : Token { 
  public readonly char sep;
  
  public Sep(char sep) {
    this.sep = sep;
  }

  public override String ToString() {
    return String.Format("sep:{0}", sep);
  }
}
