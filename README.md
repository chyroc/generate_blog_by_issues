# generate_blog_by_issues

[![Build Status](https://travis-ci.org/Chyroc/generate_blog_by_issues.svg?branch=master)](https://travis-ci.org/Chyroc/generate_blog_by_issues)

# install
```
go get -u github.com/chyroc/generate_blog_by_issues
```

# use
```
generate_blog_by_issues -repo chyroc/chyroc.github.io -token <github_personal_token> -config <config file>
```

exampple config json:

```json
{
    "name": "Chyroc的博客",
    "host": "blog.chyroc.cn",
    "author": "Chyroc",
    "notes": [
        {
            "repo": "Chyroc/chyroc.github.io",
            "paths": [
                "_md/subslice-grow.md"
            ]
        }
    ],
    "blogrolls": [
        {
            "name": "白鹤",
            "url": "https://zhenghe-md.github.io/"
        }
    ]
}
```