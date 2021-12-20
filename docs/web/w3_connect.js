/* Moralis init code */
serverUrl = "https://6xkzapnzts81.usemoralis.com:2053/server";
appId = "Reph8y9LYJlQbTlyEWLd5iEShn5iKMLdr6uzyfc7";
Moralis.start({ serverUrl, appId });

/* Authentication code */
async function login() {
    let user = Moralis.User.current();
    if (!user) {
        user = await Moralis.authenticate({ signingMessage: "Log in using Moralis" })
            .then(function (user) {
                console.log("logged in user:", user);
                console.log(user.get("ethAddress"));
            })
            .catch(function (error) {
                console.log(error);
            });
    }
}

async function logOut() {
    await Moralis.User.logOut();
    console.log("logged out");
}

document.getElementById("moralis-connector-button").onclick = login;
//document.getElementById("btn-logout").onclick = logOut;