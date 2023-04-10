
> A Powerful Bot to Update GitHub Profile's Readme using GitHub Actions

## Usage

1. Go to your repository
2. Add the following to your **README.md** file, you can use any title. Just make sure that you use 
`<!-- BLOG-POST-LIST:START --><!-- BLOG-POST-LIST:END -->` in your readme. The workflow will replace this comment with the actual blog posts list:

    ```markdown
    # Blog posts
    <!-- BLOG-POST-LIST:START -->
    <!-- BLOG-POST-LIST:END -->
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
        uses: armingli/readme-update-actions@v2
        with:
          rss_list: "https://blog.metaprogramming.space/rss.xml" # required
          max_post: 3  # optional, default 5 
```

5. Replace the above URL list with your own RSS feed URLs.
6. Commit and wait for it to run automatically, or you can also trigger it manually to see the result instantly.


<!-- BLOG-POST-LIST:START -->
<!-- BLOG-POST-LIST:END -->