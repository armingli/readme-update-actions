![readme-update-actions](./public/images/brand-name.png)

> A Powerful Bot to Update GitHub Profile's Readme using GitHub Actions

<p align="center">
    <img src="public/images/brand.jpg" height="150"><br>
    <a href="https://github.com/imskr/readme-update-actions/releases"><img alt="GitHub release (latest by date including pre-releases)" src="https://img.shields.io/github/v/release/imskr/readme-update-actions?include_prereleases&style=flat-square"></a>
    <a href="https://github.com/imskr/readme-update-actions/actions/workflows/build.yml"><img alt="Actions workflow" src="https://img.shields.io/github/workflow/status/imskr/readme-update-actions/Build/main?style=flat-square"></a>
    <a href="https://github.com/imskr/readme-update-actions/issues"><img alt="Github Issues" src="https://img.shields.io/github/issues/imskr/readme-update-actions?color=orange&style=flat-square"></a>
</p>
<hr noshade>

## Usage

1. Go to your repository
2. Add the following to your **README.md** file, you can use any title. Just make sure that you use `<!-- BLOG-LIST-START --><!-- BLOG-LIST-END -->` in your readme. The workflow will replace this comment with the actual blog posts list:

    ```markdown
    # Blog posts
    <!-- BLOG-LIST-START -->
    <!-- BLOG-LIST-END -->
    ```

3. Create a folder `.github/workflows` inside root of the repository if it doesn't exists.
4. Create a new file `readme-update-actions.yml`  inside `.github/workflows/`  with the following contents:

```
name: Readme Update Blog
on:
  schedule: # Run workflow automatically
    - cron: '0 * * * *' # Runs every hour
  workflow_dispatch: # Run workflow manually (without waiting for the cron to be called), through the GitHub Actions Workflow page directly

jobs:
  update-readme-with-blog:
    name: Update latest blogs
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3
      - name: Fetch posts
        uses: imskr/readme-update-actions@v2
        with:
          RSS_LIST: "https://blog.metaprogramming.space/rss.xml" # required
          MAX_POST: 5  # optional, default 3
          COMMIT_USER: "commiter"  # optional
          COMMIT_EMAIL: "someone@example.com" # optional
          COMMIT_MESSAGE: "Update readme with latest blogs" # optional, offer default msg like this
```

5. Replace the above URL list with your own RSS feed URLs.
6. Commit and wait for it to run automatically, or you can also trigger it manually to see the result instantly.


<!-- BLOG-LIST-START -->
<!-- BLOG-LIST-END -->

