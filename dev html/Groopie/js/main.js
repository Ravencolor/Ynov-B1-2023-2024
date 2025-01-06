// Ecouteur d'événements pour le formulaire de recherche
document.getElementById('searchForm').addEventListener('submit', function(event) {
    // Empêcher l'envoi du formulaire par défaut
    event.preventDefault();

    var query = document.getElementById("searchInput").value;

    fetch("/search?q=" + query)
        .then(response => response.json())
        .then(data => {
            // Traitez les résultats de la recherche ici
            console.log("Résultats de la recherche :", data);
            var searchResults = document.getElementById("searchResults");
            searchResults.innerHTML = ""; // Efface les résultats précédents
            data.forEach(result => {
                var link = document.createElement("a");
                link.href = result.url;
                link.textContent = result.title;
                var listItem = document.createElement("li");
                listItem.appendChild(link);
                searchResults.appendChild(listItem);
            });
        })
        .catch(error => {
            // Gérer les erreurs en cas de problème lors de la recherche
            console.error('Erreur lors de la recherche:', error);
        });
});


// Attendre le chargement complet du DOM avant de récupérer la liste des groupes
document.addEventListener('DOMContentLoaded', function() {
    var searchInput = document.getElementById('searchInput');
    var groupList = document.querySelectorAll('.flip-card');

    if (searchInput && groupList.length > 0) {
        searchInput.addEventListener('input', function() {
            var searchText = searchInput.value.toLowerCase();

            groupList.forEach(function(group) {
                var groupName = group.querySelector('input[type="submit"]').value.toLowerCase();

                if (groupName.includes(searchText)) {
                    group.style.display = 'block';
                } else {
                    group.style.display = 'none';
                }
            });
        });
    } else {
        console.error('Les éléments requis ne sont pas trouvés dans le DOM.');
    }
});

function search() {
    var searchInput = document.getElementById("searchInput").value;

    // Envoyer une requête AJAX au serveur avec le terme de recherche
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/search?q=" + encodeURIComponent(searchInput), true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState == XMLHttpRequest.DONE) {
            if (xhr.status == 200) {
                var response = JSON.parse(xhr.responseText);
                displaySearchResults(response); // Appeler la fonction pour afficher les résultats
            } else {
                console.error("Une erreur s'est produite lors de la recherche:", xhr.status);
            }
        }
    };
    xhr.send();
}

document.getElementById('searchForm').addEventListener('submit', function(event) {
    // Empêcher l'envoi du formulaire par défaut
    event.preventDefault();

    var query = document.getElementById("searchInput").value;

    fetch("/search?q=" + query)
        .then(response => response.json())
        .then(data => {
            // Traitez les résultats de la recherche ici
            console.log("Résultats de la recherche :", data);
            var searchResults = document.getElementById("searchResults");
            searchResults.innerHTML = ""; // Efface les résultats précédents
            data.forEach(result => {
                var flipCard = document.createElement("div");
                flipCard.className = "flip-card";

                var flipCardInner = document.createElement("div");
                flipCardInner.className = "flip-card-inner";

                var flipCardFront = document.createElement("div");
                flipCardFront.className = "flip-card-front";

                var artistNameInput = document.createElement("input");
                artistNameInput.type = "hidden";
                artistNameInput.name = "ArtistName";
                artistNameInput.value = result.Name;
                artistNameInput.alt = result.Name;

                var artistImageInput = document.createElement("input");
                artistImageInput.type = "Image";
                artistImageInput.src = result.Image;

                flipCardFront.appendChild(artistNameInput);
                flipCardFront.appendChild(artistImageInput);

                var flipCardBack = document.createElement("div");
                flipCardBack.className = "flip-card-back";

                var artistNameButton = document.createElement("input");
                artistNameButton.type = "submit";
                artistNameButton.value = result.Name;

                flipCardBack.appendChild(artistNameButton);

                flipCardInner.appendChild(flipCardFront);
                flipCardInner.appendChild(flipCardBack);

                flipCard.appendChild(flipCardInner);

                searchResults.appendChild(flipCard);
            });
        })
        .catch(error => {
            // Gérer les erreurs en cas de problème lors de la recherche
            console.error('Erreur lors de la recherche:', error);
        });
});