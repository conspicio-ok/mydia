const cookie = {}


cookie.get = function(name)
{
	return document.cookie
		.split("; ")
		.find(row => row.startsWith(name + "="))
		?.split("=")[1];
}

const token = cookie.get("token");
if (token)
	console.log(token)
else
	document.cookie = "token=abc123; path=/; max-age=3600"; // expire dans 1h
