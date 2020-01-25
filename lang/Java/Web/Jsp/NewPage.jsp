<Html><Head></Head><Body>
<%@ page session="true" %>
The Current Session id is: <%= session.getId() %>
 
 
Checking the value stored in our validate attribute:
<%=session.getAttribute("Validate")%>
</Body></Html>
                 
