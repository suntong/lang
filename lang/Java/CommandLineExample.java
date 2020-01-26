class CommandLineExample{  
    public static void main(String args[]){
	ex11(args);
	ex12(args);
	// ex2
	ex21(args);
	ex22(args);
    }

    // == ex1 - https://www.javatpoint.com/command-line-argument
    
    public static void ex11(String args[]){  
	System.out.println("Your first argument is: "+args[0]);  
    }

    // also https://www.studytonight.com/java/command-line-argument.php
    public static void ex12(String args[]){  
	for(int i=0;i<args.length;i++)  
	    System.out.println(args[i]);  
    }
    
    // == ex2 - https://docs.oracle.com/javase/tutorial/essential/environment/cmdLineArgs.html
    
    public static void ex21(String args[]){
        for (String s: args) {
            System.out.println(s);
        }
    }

    /*
      If an application needs to support a numeric command-line argument, it
      must convert a String argument that represents a number, such as "34",
      to a numeric value. Here is a code snippet that converts a
      command-line argument to an int:
    */
    public static void ex22(String args[]){
	int firstArg;
	if (args.length > 0) {
	    try {
		firstArg = Integer.parseInt(args[0]);
	    } catch (NumberFormatException e) {
		System.err.println("Argument" + args[0] + " must be an integer.");
		System.exit(1);
	    }
	}
    }
}

