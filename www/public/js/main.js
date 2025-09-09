const cookie = {}
const api = {}


cookie.get = function(name)
{
	return document.cookie
		.split("; ")
		.find(row => row.startsWith(name + "="))
		?.split("=")[1];
}

api.url = "https://mydia.com/api"

api.getUser = async function () {
	try
	{
		const res = await fetch(api.url + "/users")
		const data = await res.json()
		return data
	}
	catch (err)
	{
		console.error("Erreur fetch users:", err)
	}
}