# Blog API

Simple Go BE to serve blog posts

## Authors

- [@klaborda](https://www.github.com/klaborda)

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`PSQL_HOSTNAME`
`PSQL_USERNAME`
`PSQL_PASSWORD`

## API Reference

#### Get all items

```http
  GET /posts
```

#### Get a post by id

```http
  GET /posts/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of post to fetch |

#### Update a post by id

```http
  PATCH /posts/:id
```

| Parameter | Type     | Description                        |
| :-------- | :------- | :--------------------------------- |
| `id`      | `string` | **Required**. Id of post to update |

#### Delete a post by id

```http
  DELETE  /posts/:id
```

| Parameter | Type     | Description                        |
| :-------- | :------- | :--------------------------------- |
| `id`      | `string` | **Required**. Id of post to delete |
