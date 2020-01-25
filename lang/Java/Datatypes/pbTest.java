import java.io.FileReader;
import java.io.BufferedReader;
import java.io.IOException;

import java.util.Arrays;

class PhoneBook   
{
  static final int MAX_ENTRIES = 1000;
  
  public String name[] = null;
  public String phone[] = null;
  public String nameNdx[] = null;
  public String phoneNdx[] = null;
  public int entryMax = 0;
  
  public void readFile(String fname)
  {
	FileReader reader = null;
	BufferedReader in = null;
      try
      {
		reader = new FileReader(fname);
		in = new BufferedReader(reader);
      }
      
	  catch(IOException e)
      {  
		  System.out.println("Error opening file");
          System.exit(1);
      }

      
      try
      {  
		  processFile(in);
          in.close();
      }
      
	  catch(IOException e)
      {  
		  System.out.println("Error processing file");
          System.exit(1);
      }
    }
  
   /**
      Encrypts all characters in a file.
      @param in the plaintext file
   */      
   public void processFile(BufferedReader in) throws IOException
   {
	     String tname[] = new String[MAX_ENTRIES];
	     String tphone[] = new String[MAX_ENTRIES];
	     String inputLine;
	     
		 while(true)
	     {
			 inputLine = in.readLine();
			 if (inputLine==null) break;
			 tname[entryMax] = inputLine;
			 inputLine = in.readLine();
			 tphone[entryMax] = inputLine;
			 entryMax++;
		 }
			
	     name =  new String[entryMax];
	     phone = new String[entryMax];
	     nameNdx =  new String[entryMax];
	     phoneNdx = new String[entryMax];

	     System.arraycopy((Object)tname, 0, (Object)name, 0, entryMax);
	     System.arraycopy((Object)tphone, 0, (Object)phone, 0, entryMax);

	//       sort names and phone#
	//       java.util.Arrays.sort((Object[])name);
	//       java.util.Arrays.sort((Object[])phone);
	    
		 StringSort.sort(name);
	     StringSort.sort(phone);

	     // fill lookup table
	     for (int i = 0; i < entryMax; i++)
		 {
	      //System.out.println(i + ": " + tname[i] + ", " + search(name, tname[i]));
	        phoneNdx[search(name, tname[i])] = tphone[i];
	     }
	   
	     for (int i = 0; i < entryMax; i++)
	     {
	       nameNdx[search(phone, tphone[i])] = tname[i];
	     }
     
     //print(phone);
     }
   
  
   /**
      Finds a value in a range of a sorted array, 
      using the binary search algorithm.
      @param stra the sorted array
      @param from the first index in the range to search
      @param to the last index in the range to search
      @param v the value to search
      @return the index at which the value occurs, or -1
      if it does not occur in the array
   */
	   public static int binarySearch(String[] stra, 
	      int from, int to, String v)
	   {  if (from > to)
	         return -1;
	         
	      int mid = (from + to) / 2;
	      int diff = stra[mid].compareToIgnoreCase(v);
		
	      if (diff == 0) // stra[mid] == v
	         return mid;
			 //?
	      else if (diff < 0) // stra[mid] < v 
	         return binarySearch(stra, mid + 1, to, v);
	      else
	         return binarySearch(stra, from, mid - 1, v);
	   }
	   
   /**
      Finds a value in a sorted array, using the binary
      search algorithm.
      @param stra the sorted array
      @param v the value to search
      @return the index at which the value occurs, or -1
      if it does not occur in the array
   */
	   public static int search(String[] stra, String v)
	   {  return binarySearch(stra, 0, stra.length - 1, v);
	   }

  /** 
      Prints all elements in an array.
      @param a the array to print
   */
	   public static void print(String[] a)
	   {  for (int i = 0; i < a.length; i++)
	         System.out.println(a[i]);
	   }
	
	public void nsNumber()
	{
		try
		{
		   ConsoleReader console = new ConsoleReader(System.in);
		   System.out.println("Enter name to search phone number: ");
		   String m = console.readLine();
		   int k = search(name, m);
		   if (k < 0)
		   {
		   	System.out.println("No such a record.");
		   }
		   
		   else
		   {
			   System.out.println("Found in position " + k);
			   System.out.println("Phone: " + phoneNdx[k]);
			   System.out.println("Name: " + name[k]);
		   }
		   
		
		}
		catch(NumberFormatException e)
		{
			System.out.println("Input error" + e);
		}

	}

	public void nsName()
	{
		try
		{
	        ConsoleReader console = new ConsoleReader(System.in);
	        System.out.println("Enter phone number to search name:");
	        String n = console.readLine();
        	int j = search(phone, n);
			if (j < 0)
			{
				System.out.println("No such a record.");
			}
			else
			{
	        	System.out.println("Found in position " + j);
		        System.out.println("Name: " + nameNdx[j]);
		        System.out.println("Phone: " + phone[j]);
			}
		
		}
		catch(NumberFormatException e)
		{
			System.out.println("Input error" + e);
		}
	}

	  
	public void run()
	{
		  try
	      {
	       int x;
	       
	       do								 
	       {
		   	System.out.println("");
	       	System.out.println("Please enter the following option.");
	       	System.out.println("1: enter name to search number.");
	       	System.out.println("2: enter number to serach name.");
	       	System.out.println("0: to quit.");
			System.out.println("");
	       	
	       	ConsoleReader console = new ConsoleReader(System.in);
	       	x = console.readInt();
	       	switch (x)
	      	{
	      	 	case 1: nsNumber(); break;
	      		case 2: nsName(); break;
	      		case 0: break;
	      	}
	       }while (x != 0);
	      }
		  
	      catch(NumberFormatException e)
	      {
	      	System.out.println("Input error" + e);
	      }
		  
	}
	
	

}

class StringSort
{
  /**
      Sorts an array.
      @param a the array to sort
   */
   
   //?
   public static void sort(String[] a)
   {  for (int n = 0; n < a.length - 1; n++)
      {  int minPos = minimumPosition(a, n);
         if (minPos != n)
            swap(a, minPos, n);
      }
   }

  /**
      Finds the smallest element in an array range.
      @param a the array to search
      @param from the first position in a to compare
      @return the position of the smallest element in the
      range a[from]...a[a.length - 1]
   */
   public static int minimumPosition(String[] a, int from)
   {  int minPos = from;
      for (int i = from + 1; i < a.length; i++)
	  	//?
         if (a[i].compareToIgnoreCase(a[minPos]) < 0) minPos = i;
      return minPos;
   }

  /**
      Swaps two elements in an array.
      @param a the array with the elements to swap
      @param i the index of one of the elements
      @param j the index of the other element
      @return
   */
   public static void swap(String[] a, int i, int j)
   {  String temp = a[i];
      a[i] = a[j];
      a[j] = temp;
   }
   
}


	  
	  
public class pbTest
{  public static void main(String[] args)
   {  
      if (args.length != 1) usage();

      PhoneBook pb = new PhoneBook();
      // gather command line arguments and open files
      pb.readFile(args[0]);
	  pb.run();
	  
   }
   
  /**
      Prints a message describing proper usage and exits.
   */
   public static void usage()
   {  System.out.println
         ("Usage: java pbTest infile");
      System.exit(1);
   }

}

