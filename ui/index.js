const form = document.getElementById("new-url-form");
        const content = document.getElementById("new-url-content");
        form.addEventListener("submit", (e) => {
            e.preventDefault();
            const url = form.url.value;
            fetch("http://localhost:8080/api/shorten", {
            method: "POST",
            body: JSON.stringify({
                url: url,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8",
            },
            })
            .then((response) => {
                return response.json();
            })
            .then((data) => {
                content.innerHTML = `Your shortened URL is: <a href="http://localhost:8080/short/${data.shortUrl}" target="_blank">http://localhost:8080/short/${data.shortUrl}</a>, <br>and your id is '${data.id}' (save this for future url deletion)`;
                form.url.value = "";
                content.style.display = "block";
            })
            .catch((err) => {
                content.innerHTML = err;
                console.log(err);
                content.style.display = "block";
            });
        });

        const deleteForm = document.getElementById("delete-url-form");
        const deleteContent = document.getElementById("delete-url-content");
        deleteForm.addEventListener("submit", (e) => {
            e.preventDefault();
            const id = deleteForm.id.value;
            fetch(`http://localhost:8080/api/delete/${id}`, {
            method: "DELETE",
            headers: {
                "Content-type": "application/json; charset=UTF-8",
            },
            })
            .then((response) => {
                if (!response.ok) {
                    throw new Error(response.status + " " + response.statusText);
                }

                deleteContent.innerHTML = `Your url with id ${id} was deleted`;
                deleteForm.id.value = "";
                deleteContent.style.display = "block";
            }).catch((err) => {
                deleteContent.innerHTML = err;
                console.log(err);
                deleteContent.style.display = "block";
            });
        })