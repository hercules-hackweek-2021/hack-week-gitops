## GitOps

Team: Andrey Garshyn, Hercules Merscher, Roshan Patil, Sahil Lone

Problem: Can we know when a service was exactly deployed to theta/stage/live? What was the version/revision deployed?

Solution: To use GoCD to bring visibility and autonomy to the teams within SumUp to come up with their own deployment pipelines.

Result: 
1. Visibility of which revision is deployed to multiple environments.
2. Single source of truth being the application repository.
3. Easily promote revisions from stage to live.
4. Easily rollback to a previous revision.

Moving forward: Install and configure GoCD in our current infrastructure, and start using it with real projects.
