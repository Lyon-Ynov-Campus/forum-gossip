
export async function toggleLike(postId) {
    try {
    
        const response = await fetch(`/api/posts/like`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ id: postId }) 
        });

        if (!response.ok) {
            throw new Error("Erreur lors de la requête");
        }

        
        const data = await response.json(); 
        

        updateLikeUI(postId, data.likesCount, data.hasLiked);

    } catch (error) {
        console.error("Erreur API Like:", error);
        alert("Impossible de liker pour le moment.");
    }
}

function updateLikeUI(postId, count, hasLiked) {
    const countElement = document.getElementById(`count-${postId}`);
    const btnElement = document.getElementById(`btn-like-${postId}`);

    if (countElement) countElement.innerText = count;
    
    if (btnElement) {
        
        btnElement.classList.toggle('is-liked', hasLiked);
        btnElement.style.color = hasLiked ? 'red' : 'black';
    }
}