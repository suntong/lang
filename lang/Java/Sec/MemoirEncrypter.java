/*
 * MemoirEncrypter - A program to write secret memoirs!
 * Copyright 1998 by Digital Focus and Clifford Berg. 
 */

import javax.crypto.*;
import javax.crypto.spec.*;
import java.awt.*;
import java.awt.event.*;
import java.io.*;

/**
 * Main class for the memo encrypter.
 */
 
public class MemoirEncrypter extends Frame {
  public static void main(String[] args) throws Exception {
    // Construct and display the UI component
    MemoirEncrypter f = new MemoirEncrypter();
    f.show();
    }
        
  public MemoirEncrypter() throws Exception {
    super("MemoirEncrypter");
    setSize(600, 300);
                
    //
    // Create some controls
    //
                
    // A button for opening a file dialog
    openButton = new Button("Open");
    openButton.addActionListener
      (
       new ActionListener() {
	   public void actionPerformed(ActionEvent e) {
	     getPath(FileDialog.LOAD);
	     if (path == null) return;
                                        
	     if (password == null) getPassword();
	     if (password == null) return;
                                        
	     // Compute the cipher
	     Cipher cipher = null;
	     try {
		 cipher = computePBECipher(password, salt, 
					   iterations, cipherName, Cipher.DECRYPT_MODE);
		 }
	     catch (Exception ex) { ex.printStackTrace(); return; }
                                        
	     FileInputStream fis = null;
	     try {
		 // Open the file
		 fis = new FileInputStream(path);
		 CipherInputStream cis = new CipherInputStream(fis, cipher);
		 InputStreamReader isr = new InputStreamReader(cis);
		 BufferedReader br = new BufferedReader(isr);
                                        
		 // Read in the file, and decrypt it, a line at a time
		 textArea.setText("");
		 for (;;) {
		     String line = br.readLine();
		     if (line == null) break;
		     textArea.append(line + "\n");
		     }
                                                
		 br.close();
		 }
	     catch (Exception ex) { ex.printStackTrace(); return; }
	     }
	   }
       );
                
    // A button for saving
    saveButton = new Button("Save");
    saveButton.addActionListener
      (
       new ActionListener() {
	   public void actionPerformed(ActionEvent e) {
	     if (path == null) getPath(FileDialog.SAVE);
	     if (path == null) return;

	     if (password == null) getPassword();
	     if (password == null) return;
                                        
	     // Compute the cipher
	     Cipher cipher = null;
	     try {
		 cipher = computePBECipher(password, salt, 
					   iterations, cipherName, Cipher.ENCRYPT_MODE);
		 }
	     catch (Exception ex) { ex.printStackTrace(); return; }

	     // Open the file
	     FileOutputStream fos = null;
	     try {
		 fos = new FileOutputStream(path);
		 CipherOutputStream cos = new CipherOutputStream(fos, cipher);
                                        
		 // Write out the file, and encrypt it, a line at a time
		 String text = textArea.getText();
		 byte[] buf = text.getBytes();

		 cos.write(buf);
		 cos.flush();
		 cos.close();
		 }
	     catch (Exception ex) { ex.printStackTrace(); return; }
	     }
	   }
       );
                
    //
    // Create an area to display the text
    //
                
    textArea = new TextArea();
                
    //
    // Add all these components, using an appropriate layout
    //
                
    controlPanel = new Panel();
    controlPanel.add(openButton);
    controlPanel.add(saveButton);
    add("North", controlPanel);
    add("Center", textArea);
                
    validate();
    }

  protected void getPath(int mode) {
    FileDialog fd = 
      new FileDialog(this, 
		     "Select a file to open", mode);
    fd.show();
    String d = fd.getDirectory();
    if (d == null) return;
    System.out.println("d=" + d);
    String f = fd.getFile();
    if (f == null) return;
    System.out.println("f=" + f);
    path = d + f;
    }

  protected void getPassword() {
    PasswordDialog pd =
      new PasswordDialog(MemoirEncrypter.this, 
			 "Enter password", true);
    pd.show();
    password = pd.getPassword();
    }

  protected static Cipher computePBECipher
    (String password, byte[] salt, int iterations, String cipherName, int mode)
    throws Exception {
    // Compute the key
    PBEParameterSpec pbeParamSpec = new PBEParameterSpec(salt, iterations);
    PBEKeySpec pbeKeySpec = new PBEKeySpec(password.toCharArray());
    SecretKeyFactory keyFac = SecretKeyFactory.getInstance(cipherName);
    SecretKey key = keyFac.generateSecret(pbeKeySpec);

    // Construct the cipher
    Cipher cipher = Cipher.getInstance(cipherName);
    cipher.init(mode, key, pbeParamSpec);
    return cipher;
    }
        
  private Button openButton;
  private Button saveButton;
  private String password;
  private Panel controlPanel;
  private TextArea textArea;
  private String path;

  private final byte[] salt = { 
    (byte)0xaa, (byte)0xbb, (byte)0xcc, (byte)0xdd,
    (byte)0x22, (byte)0x44, (byte)0xab, (byte)0x12 };
  private final int iterations = 10;
  private final String cipherName = "PBEWithMD5AndDES";

  static class PasswordDialog extends Dialog {
    public PasswordDialog(Frame parent, String title, boolean modal) {
      super(parent, title, modal);
      setLayout(new FlowLayout());
      setSize(300, 80);
                        
      textField = new TextField(10);
      textField.setEchoChar('*');
      textField.setBackground(Color.white);
      add(textField);
                        
      closeButton = new Button("OK");
      closeButton.addActionListener
	(
	 new ActionListener() {
	     public void actionPerformed(ActionEvent e) {
	       // User has completed entering a password
	       pswd = textField.getText();
	       dispose();
	       return;
	       }
	     }
	 );
      add(closeButton);
      }
                
    public String getPassword() { return pswd; }
                
    private TextField textField;
    private Button closeButton;
    private String pswd;
    }
  }

