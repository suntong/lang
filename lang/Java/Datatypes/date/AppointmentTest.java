import java.util.StringTokenizer;
import java.util.Vector;
import java.util.Calendar;
import java.io.IOException;
import java.text.DateFormat;


public class AppointmentTest {
	
  public AppointmentTest() {
    appointments = new Vector();
    }
	
  public void addAppointment(Appointment anAppointment) {
    appointments.add(anAppointment);
    }
	
  public void removeAppointment() {
    try {
      ConsoleReader console = new ConsoleReader(System.in);
      System.out.println("Date of the canceled appointment.");
      String SDateOfCan = console.readLine();
      Calendar CDateOfCan = parseDate(SDateOfCan);
      }
    catch(NumberFormatException e) {
      System.out.println("Input error" + e);
      }
		
		
    System.out.println("Appointments are: " );
    for(int i = 0; i<appointments.size();i++) {
      if ((((Appointment)appointments.get(i)).getDate()).equals(FixDate)) {
	System.out.println(i + appointments.get(i).toString());
	}
      }
		
    try {
      ConsoleReader console = new ConsoleReader(System.in);
      System.out.println("Enter the number of the appointment.");
      int number = console.readInt();
      appointments.remove(appointments.get(number));
      }
    catch(NumberFormatException e) {
      System.out.println("Input error" + e);
      }
		
    }
	
  public Calendar parseDate(String date) {
    StringTokenizer token = new StringTokenizer(date, "/");
    int theYear = Integer.parseInt(token.nextToken());
    int theMonth = Integer.parseInt(token.nextToken());
    int theDate = Integer.parseInt(token.nextToken());
		
    System.out.println(theYear);
    System.out.println(theMonth);
    System.out.println(theDate);
		
    //		while (token.hasMoreTokens())
    //		{
    //			String token = token.nextToken();
    //			
    //		}
		
    Calendar dateCalendar = Calendar.getInstance();
    try {
      dateCalendar.set(theYear, theMonth, theDate);
      }	 
		
    catch(NumberFormatException e) {
      System.out.println("Input error" + e);
      }
    System.out.println(dateCalendar.toString());
    return dateCalendar;
    }
	
  public Calendar parseTime(Calendar date,String time) {
		
		
    Calendar timeCalendar = Calendar.getInstance();
    try {
      StringTokenizer token = new StringTokenizer(time, ":");
      timeCalendar.set(date.YEAR, date.MONTH, date.DATE,
		       Integer.parseInt(token.nextToken()), 
		       Integer.parseInt(token.nextToken()));
      //			 System.out.println();
      //			 System.out.println();
      //				 
		 
			
      }
		
		
    catch(NumberFormatException e) {
      System.out.println("Input error" + e);
      }
		
    return timeCalendar;
			
    }
		
	
  public void printConflict() {
    for(int i = 0; i<appointments.size();i++) {
      for(int j = i + 1; j<appointments.size(); j++) {
	if (((Appointment)appointments.get(i)).isConflict(((Appointment)appointments.get(j)))) {
	  System.out.println("Found the following conflict" + 
			     appointments.get(i).toString() + 
			     appointments.get(j).toString());
	  }
	}
      }
    }
	
	
  public void printADay() {
    try {
      ConsoleReader console = new ConsoleReader(System.in);
      System.out.println("Please enter a date. year/month/date");
      String pDate = console.readLine();
      Calendar FixDate = parseDate(pDate);
			
      System.out.println("Appointments are: " );
      for(int i = 0; i<appointments.size();i++) {
	if ((((Appointment)appointments.get(i)).getDate()).equals(FixDate)) {
	  System.out.println(appointments.get(i).toString());
	  }
	}
      }
		
    catch(NumberFormatException e) {
      System.out.println("Input error" + e);
      }
    }
	
	
  public void printAll() {
    for(int i = 0; i<appointments.size(); i++) {
      Appointment exAppointment = ((Appointment)appointments.get(i));
      String sExAppointment = exAppointment.toString();
      System.out.println(sExAppointment);
      }
    }
	
  public void inputAppointment() {
    try {
      ConsoleReader console = new ConsoleReader(System.in);
			
      System.out.println("Please enter the description.");
      String strDes = console.readLine();
				
			
      System.out.println("Please enter date. YEAR/MONTH/DATE");
      String strDate = console.readLine();
      Calendar date = parseDate(strDate);
						
      System.out.println("Please enter starting time. HOUR:MINUT");
      String strStartTime = console.readLine();
      Calendar startTime = parseTime(date, strStartTime);
			
      System.out.println("Please enter ending time. HOUR:MINUT");
      String strEndTime = console.readLine();
      Calendar endTime = parseTime(date, strEndTime);
			
      Appointment sampleAppointment = new Appointment(strDes ,date, startTime, endTime);
      appointments.add(sampleAppointment);
      }
		
    catch(NumberFormatException e) {
      System.out.println("Input error" + e);
      }
    }
	
  public void run() {
    FixDate = Calendar.getInstance();
    FixDate.set(2001,11,30);
    Appointment a = new Appointment();
    System.out.println(a.date(FixDate));
		
		
    try {
      int x;
			 
      do {
	System.out.println("\nPlease choose the following option:");
	System.out.println("1: add an appointment");
	System.out.println("2: remove an appointment");
	System.out.println("3: find appointments in a particular day");
	System.out.println("4: check if to appointments conflict");
	System.out.println("5: print all appointments");
	System.out.println("===========");
	System.out.println("0: to quit\n");

	ConsoleReader console = new ConsoleReader(System.in);
	x = console.readInt();
	switch (x) {
	case 1: inputAppointment(); break;
	case 2: removeAppointment(); break;
	case 3: printADay(); break;
	case 4: printConflict(); break;
	case 5: printAll(); break;
	case 0: break;
	  }
	}while (x != 0);
      }
		 
    catch(NumberFormatException e) {
      System.out.println("Input error" + e);
      }
    }
	
  private Vector appointments;
  private Calendar FixDate;



  private class Appointment {
    public Appointment() {
      des = null;
      date = null;
      startTime = null;
      endTime = null;
      }
		
    public String toString() {
      //			DateFormat df = DateFormat.getDateInstance();
      //			return des + " "+ df.format(date.getTime()) +" " +startTime.toString() + " "+endTime.toString();
      return des +" "+ date(date) +" "+ time(startTime);
      }
			
    public String date(Calendar adate) {
      //System.out.println(adate.toString());
      String aday = adate.get(Calendar.YEAR) + "/" +
	adate.get(Calendar.MONTH) + "/" +
	adate.get(Calendar.DAY_OF_MONTH);
      return aday;
			
      }
		
    public String time(Calendar time) {
      return time.get(Calendar.HOUR) + ":" + time.get(Calendar.MINUTE);
      }
		
    public Appointment(String aDes, Calendar aDate, Calendar aStartTime, Calendar aEndTime) {
      des = aDes;
      date = aDate;
      startTime = aStartTime;
      endTime = aEndTime;
      }
		
    public Calendar getDate() {
      return date;
      }
		
    public boolean isConflict(Appointment anAppointment) {
      boolean isConflict = false;
      //this
      if (startTime.after(anAppointment.startTime))
	return 
	  startTime.before(anAppointment.endTime);
      else
	return 
	  endTime.after(anAppointment.startTime);
			
      }
		
    private String des;
    private Calendar date;
    public Calendar startTime;
    public Calendar endTime;
		
    }
	
  public static void main(String[] args) {
    AppointmentTest aCalendar = new AppointmentTest();
    aCalendar.run();
    }
  }
