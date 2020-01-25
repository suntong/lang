import java.util.*;
import java.text.*;

public class DateTest {

  static private void test(){
    String dateR = "1998.12.31";
    SimpleDateFormat formatter
      = new SimpleDateFormat("yyyy.MM.dd");
    ParsePosition pos = new ParsePosition(0);
    Date date = formatter.parse(dateR, pos);
    System.out.println(date);
    }

  public static void main(String[] args) {

// ............................................................. &ntc ...
    System.out.println("\n*** New test case ***\n");

    // Format the current time.
    SimpleDateFormat formatter
      = new SimpleDateFormat ("yyyy.MM.dd G 'at' hh:mm:ss a zzz");
    Date currentTime_1 = new Date();
    String dateString = formatter.format(currentTime_1);
    System.out.print("Current time is: ");
    System.out.println(dateString);

  // Parse the previous string back into a Date.
    ParsePosition pos = new ParsePosition(0);
    Date currentTime_2 = formatter.parse(dateString, pos);
    System.out.print("Parsed time is: ");
    System.out.println(currentTime_2);

    SimpleDateFormat timeFormatter
      = new SimpleDateFormat("yyyy.MM.dd-hha");
    dateString = timeFormatter.format(currentTime_1);
    System.out.print("Current time is: ");
    System.out.println(dateString);

    pos = new ParsePosition(0);
    currentTime_2 = timeFormatter.parse(dateString, pos);
    System.out.print("Parsed time is: ");
    System.out.println(currentTime_2);

// ............................................................. &ntc ...
    System.out.println("\n*** New test case ***\n");

    // get the current time in milliseconds 
    long time = System.currentTimeMillis();

    // provide to user as the default
    formatter.applyPattern("yyyy-dd-MM HH:mm:ss");
    dateString = new String(formatter.format(new Date(time)));
    System.out.print("Current time is: ");
    System.out.println(dateString);

    // read it back from user input
    pos = new ParsePosition(0);	// very important!
    time = formatter.parse(dateString, pos).getTime();
    System.out.print("Parsed time is: ");
    System.out.println(time);

// ............................................................. &ntc ...
    System.out.println("\n*** New test case ***\n");

    // Initialize time fields with the current date and time
    Calendar rightNow = Calendar.getInstance();
    System.out.print("Current time is: ");
    System.out.println(rightNow.getTime());
    System.out.println(rightNow);

    // coin a new one
    Calendar coined = Calendar.getInstance();
    coined.set(1998, 11, 31);
    System.out.print("Coined time is: ");
    System.out.println(coined.getTime());
    System.out.println(coined);

    Calendar coined2 = coined;
    coined2.set(Calendar.MONTH, 0);
    System.out.print("\nCoined2 time is: ");
    System.out.println(coined2.getTime());


// ............................................................. &ntc ...
    System.out.println("\n*** New test case ***\n");

    test();
    }
}
