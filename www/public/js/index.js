let users = 0


async function init()
{
	users = (await api.getUser())?.length || 0;
}

async function main()
{
	await init();
	if (users)
	{
		document.getElementById('rpassword-form').style.display = 'none';
	}
	else
	{
		let msg = document.getElementById('warning-message');
		msg.innerHTML = '<p>No user found, please create new</p>';
		msg.style.display = 'block';
		document.getElementById("form").addEventListener('submit', async function(event) {
			event.preventDefault();
			const username = document.getElementById("username").value
			const password = document.getElementById("password").value
			const rpassword = document.getElementById("rpassword").value
			if (password == rpassword)
				console.log(await api.login(username, password));
		});
	}

	// const userJWT = cookie.get("userJWT");
	// if (userJWT)
	// 	console.log(userJWT)
}

main()
