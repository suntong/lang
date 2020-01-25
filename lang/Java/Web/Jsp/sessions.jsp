
<%@ page language=javascript %><%
/* session.jsp
*
* Trivial example of session variables, just stores a counter
* in the session variable.
*
* Since the Request object comes from the Servlet API, Resin
* scripts can share session variables with Java Servlets.
*/
var count = session.value.counter++;

%>

<html>
<head><title>Counter</title></head>
<body bgcolor=#ffffff>
<%

if (session.isNew()) {
%>
<h1>Welcome to a new session.</h1>
<%
} else if (session.value.name) {
%><h1>Welcome back: <%= session.value.name %> <%= count %> </h1><%
} else {
%><h1>Welcome back: <%= count %> </h1><%
}

%>
<p/><a href="<%= response.encodeUrl("session.jsp") %>">click</a> to enable
session rewriting.
</body>
</html>
