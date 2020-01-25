import javax.swing.JFrame;
import javax.swing.JPanel;
import java.awt.event.WindowAdapter;
import java.awt.event.WindowEvent;
import java.awt.geom.Ellipse2D;
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.Container;
import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.geom.Line2D;
import java.awt.geom.Point2D;

public class LineLength
{
  public static void main(String[] args)
  {
    ProgramFrame frame = new ProgramFrame();
    frame.setTitle("LineMaker");
    frame.show();
    }
  }

class ProgramFrame extends JFrame
{
  public ProgramFrame()
  {
    setSize(300, 300);
		
    WindowCloser listener = new WindowCloser();
    addWindowListener(listener);
		
    LPanel= new LinePanel();
		
    Container contentPane = getContentPane();
    contentPane.add(LPanel,"Center");
    clicks = 0;
    points = new Point2D.Double[MAX_CLICKS];
		
    MouseClickedListener Mlistener = new MouseClickedListener();
    addMouseListener(Mlistener);
		
    }
	
  private LinePanel LPanel;
  private Ellipse2D.Double mark;
  private int clicks; 
  private Point2D.Double[] points; 
  private final int MAX_CLICKS = 2;
	 
  private class LinePanel extends JPanel
  {
    public void paintComponent(Graphics g)
    {
      super.paintComponent(g);
      Graphics2D g2 = (Graphics2D)g;
			
      if (clicks == 1)
	{ 
	  double x1 = points[0].getX();
	  double y1 = points[0].getY();
	  final double RADIUS = 15;
	  Ellipse2D.Double mark
	    = new Ellipse2D.Double( x1 - RADIUS, y1 - RADIUS,
				    2 * RADIUS, 2 * RADIUS);
					 
	  g2.draw(mark);
	  }
      else if (clicks >= 2)
	{  
	  double x1 = points[0].getX();
	  double y1 = points[0].getY();
				
	  double x2 =  points[1].getX();
	  double y2 =  points[1].getY();
				
	  double midPointX = (x1+x2)/2;
	  double midPointY = (y1+y2)/2;
	  double length = Math.sqrt(Math.pow((y2- y1),2) + Math.pow(( x2- x1),2));
	  String MLength = "Length is " + length;
	  g2.drawString(MLength, (int)Math.round(midPointX), (int)Math.round(midPointY));
	  g2.draw(new Line2D.Double(points[0], points[1]));
	  }
		
      int SHeight = getHeight();
      int messageLocationY = SHeight - 10;
			
      String message = "Please click on two points.";
      g2.drawString(message,10,messageLocationY);
			
      }
    }
	
	
  private class MouseClickedListener extends MouseAdapter
  {
    public void mouseClicked(MouseEvent event)
    {  
      if (clicks >= MAX_CLICKS) 
	return;
      int mouseLocationX = event.getX();
      int mouseLocationY = event.getY();
      points[clicks] = new Point2D.Double(mouseLocationX, mouseLocationY);
      clicks++;
      repaint();
      System.out.println("asfasf");
      }
    }
	
	
  private class WindowCloser extends WindowAdapter
  {
    public void windowClosing(WindowEvent event)
    {
      System.exit(0);
      }
    }
	
  }


