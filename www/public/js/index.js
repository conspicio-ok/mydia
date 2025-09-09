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
	}

	// const userJWT = cookie.get("userJWT");
	// if (userJWT)
	// 	console.log(userJWT)
}

main()
