{{ define "posts/comment.tpl"}}
<div class="comment-card">
  <div class="comment-container">
    <h4><b>{{ .User.Username }}</b></h4>
    <p>{{ .Comment | unsafe }}</p>
  </div>
</div>
{{ end }}