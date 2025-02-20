{{ define "posts/add_edit_comment.tpl"}}
    <script type="text/javascript">
        function addOrEditComment(event, postID, user, originalComment) {
            event.preventDefault(); // Prevent form from refreshing the page
            const inputs = document.getElementsByTagName('INPUT')
            inputs[1].setAttribute('disabled', 'disabled')
            const textareas = document.getElementsByTagName('TEXTAREA')
            textareas[0].setAttribute('disabled', 'disabled')

            // get values
            comment = document.getElementById('comment').value
            csrf = document.getElementById('csrf').value

            if (originalComment) {
                // edit
                console.log("not implemented yet")
            } else {
                // new
                console.log(user)
                const newUserComment = {
                    postID: postID,
                    userID: user.id,
                    comment: comment,
                    csrf: csrf
                };
                console.log("Sending data:", JSON.stringify(newUserComment));
                fetch('/posts/' + postID + '/comments', {
                    method: 'POST',
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(newUserComment),
                })
                .then(response => {
                    if (!response.ok) {
                        document.getElementById("comment-error").style.visibility="visible";
                            throw new Error('Network response was not ok');
                        }
                        document.getElementById("comment-error").style.visibility="hidden";
                        console.log("comment added")
                        location.reload();
                    })
                .catch(error => {
                    document.getElementById("comment-error").style.visibility="visible";
                    console.error('There has been a problem with your fetch operation:', error);
                }).finally(() => {
                    inputs[1].removeAttribute('disabled')
                    textareas[0].removeAttribute('disabled')
                });

            }

        }
    </script>
    <div class="comment-card">
      <div class="comment-container">
        <form onsubmit="addOrEditComment(event, {{ .post.ID }}, {{ .user }}, {{ .comment }})">
            <label>{{ .user.Username }}'s comment<label><br>
            {{ if .comment }}
                <textarea id='comment' name='comment'>{{ .comment.Comment }}</textarea>
            {{ else }}
                <textarea placeholder='Write your comment...' id='comment' name='comment'></textarea>
            {{ end }}
            <br>
            <input type='hidden' id='csrf' name="csrf" value='{{ .csrf }}'>
            <input type='submit' value="Save">
            <div id="comment-error" style="color:red; visibility: hidden;">
                  An error occurred
            </div>
        </form>
      </div>
    </div>
{{ end }}