// Example 165 from page 133 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// 1. Run this program to see that mutual exclusion works: the sum 
//    of the bank accounts' balances forever remains 30.
// 2. Then comment out lock(this) as in 
//      /* lock (this) */ { 
//    and compile and run the program again.  Now the sum of the balances 
//    will deviate from 30 because the bank clerks occasionally overwrite 
//    each others' updates.

using System;
using System.Threading;

class Bank {
  private int account1 = 10, account2 = 20;
  public void Transfer(int amount) {
    lock (this) {
      int new1 = account1 - amount;
      Util.Pause(10);
      account1 = new1; account2 = account2 + amount;   
      Console.WriteLine("Sum is " + (account1+account2));
    }
} }

class Clerk {
  private Bank bank;
  public Clerk(Bank bank) { 
    this.bank = bank; 
  }

  public void Run() {
    for (;;) {                                  // Forever
      bank.Transfer(Util.Random(-10, 10));      //   transfer money
      Util.Pause(200, 300);                     //   then take a break
} } }

class TestBank {
  public static void Main(String[] args) {
    Bank bank = new Bank();
    Clerk clerk1 = new Clerk(bank), clerk2 = new Clerk(bank);
    new Thread(new ThreadStart(clerk1.Run)).Start();
    new Thread(new ThreadStart(clerk2.Run)).Start();
  } 
}

// Pseudo-random numbers and sleeping threads

class Util {
  private static readonly Random rnd = new Random();
  
  public static void Pause(int length) { 
    Thread.Sleep(length); 
  } 

  public static void Pause(int a, int b) { 
    Pause(Random(a, b));
  }
  
  public static int Random(int a, int b) {
    return rnd.Next(a, b);
  }
}
