import java.applet.Applet;
import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;
import java.awt.geom.Ellipse2D;

public class ClickEggApplet extends Applet
{
  public ClickEggApplet()
  {
    egg = new Ellipse2D.Double(0,0,EGG_WIDTH, EGG_HEIGHT);
    MouseClickListener listener = new MouseClickListener();
    addMouseListener(listener);
    }
	
  public void paint(Graphics g)
  {
    Graphics2D g2 = (Graphics2D)g;
    g2.draw(egg);
    }
	
  private Ellipse2D.Double egg;
  private static final double EGG_WIDTH = 30;
  private static final double EGG_HEIGHT = 30;
	
  private class MouseClickListener extends MouseAdapter
  {
    public void mouseClicked(MouseEvent event)
    {
      int mouseX = event.getX();
      int mouseY = event.getY();
      egg.setFrame(mouseX-EGG_WIDTH/2,mouseY-EGG_HEIGHT/2,EGG_WIDTH,EGG_HEIGHT);
      repaint();
      }
    }
  }
