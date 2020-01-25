import java.util.NoSuchElementException;

public class ListTest
{  public static void main(String[] args)
   {  LinkedList staff = new LinkedList();
      staff.addFirst("Sam");
      staff.addFirst("Romeo");
      staff.addFirst("Harry");
      staff.addFirst("Dick");
     
      // | in the comments indicates the iterator position
      
      LinkedList.Iterator iterator 
         = staff.listIterator(); // |DHRS
      iterator.next(); // D|HRS
      iterator.next(); // DH|RS

      // add more elements after second element
      
      iterator.add("Juliet"); // DHJ|RS
      iterator.add("Nina"); // DHJN|RS

      iterator.next(); // DHJNR|S

      // remove last traversed element 
      //iterator.remove(); // DHJN|S
      
      // print all elements
      staff.print();

      // downsize
      LinkedList.downsize(staff);
      // print all elements
      staff.print();
     }
}

/**
   A linked list is a sequence of links with efficient
   element insertion and removal.
*/

class LinkedList
{  /** 
      Constructs an empty linked list.
   */
   public LinkedList()
   {  first = null;
   }
   
   /**
      Returns the first element in the linked list.
      @return the first element in the linked list
   */
   public Object getFirst()
   {  if (first == null) 
         throw new NoSuchElementException();
      return first.data;
   }

   /**
      Removes the first element in the linked list.
      @return the removed element
   */
   public Object removeFirst()
   {  if (first == null) 
         throw new NoSuchElementException();
      Object obj = first.data;
      first = first.next;
      return obj;
   }

   /**
      Adds an element to the front of the linked list.
      @param obj the object to add
   */
   public void addFirst(Object obj)
   {  Link newLink = new Link();
      newLink.data = obj;
      newLink.next = first;
      first = newLink;
   }
   
   /**
      Returns an iterator for iterating through this list.
      @return an iterator for iterating through this list
   */
   public Iterator listIterator()
   {  return new Iterator();
   }
   
  public void print()
  {
    Iterator iterator = new Iterator();

    System.out.println("--vvvv--");
    while (iterator.hasNext())
      System.out.println(iterator.next());
    System.out.println("--^^^^--\n");
    }

  public static void downsize(LinkedList staff)
  {
    LinkedList.Iterator iterator 
      = staff.listIterator();

    while (iterator.hasNext())
      {
	iterator.next();
	iterator.next(); 

	// remove last traversed element 
	iterator.remove(); // DHJN|S
	}
    }
     
  private Link first;
   
   private class Link
   {  public Object data;
      public Link next;
   }

   public class Iterator
   {  /**
         Constructs an iterator that points to the front
         of the linked list.
      */
      public Iterator()
      {  position = null;
         previous = null;
      }
      
      /**
         Moves the iterator past the next element.
         @return the traversed element
      */
      public Object next()
      {  if (position == null)
         {  position = first;
            return getFirst();
         }
         else
         {  if (position.next == null)
               throw new NoSuchElementException();
            previous = position; // remember for remove
            position = position.next;
            return position.data;
         }
      }
      
      /**
         Tests if there is an element after the iterator 
         position.
         @return true if there is an element after the iterator 
         position
      */
      public boolean hasNext()
      {  if (position == null)
            return first != null;
         else
            return position.next != null;
      }
      
      /**
         Adds an element before the iterator position
         and moves the iterator past the inserted element.
         @param obj the object to add
      */
      public void add(Object obj)
      {  if (position == null)
            addFirst(obj);
         else
         {  Link newLink = new Link();
            newLink.data = obj;
            newLink.next = position.next;
            position.next = newLink;
            position = newLink;
            previous = null;
         }
      }
      
      /**
         Removes the last traversed element. This method may
         only be called after a call to the next() method.
      */
      public void remove()
      {  if (position == first)
            removeFirst();
         else 
         {  if (previous == null)
               throw new IllegalStateException();
            previous.next = position.next;
            position = previous;
         }
         previous = null;
      }

      private Link position;
      private Link previous;
   }
}

