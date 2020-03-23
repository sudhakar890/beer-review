# Beer review

> Simple RESTful API in Go to add and list beer reviews with data stored in MongoDB.

## Endpoints

### Get all Beer details
``` bash
GET /beers
```
### Get single Beer details
``` bash
GET /beers/{id}
```
### Get reviews of a particular Beer
``` bash
GET /beers/{id}/reviews
```

### Add Beers
``` bash
POST /beers

# Request sample
# [
# {
#    "name":"Pliny the Elder",
#    "brewery":"Russian River Brewing Company",
#    "abv":8,
#    "short_description":"Pliny the Elder is brewed with Amarillo, Centennial, CTZ, and Simcoe hops. It is well-balanced with malt, hops, and alcohol, slightly bitter with a fresh hop aroma of floral, citrus, and pine."
# },
# {
#    "name":"Oatmeal Stout",
#    "brewery":"Samuel Smith",
#    "abv":5,
#    "short_description":"Brewed with well water (the original well at the Old Brewery, sunk in 1758, is still in use, with the hard well water being drawn from 85 feet underground); fermented in ‘stone Yorkshire squares’ to create an almost opaque, wonderfully silky and smooth textured ale with a complex medium dry palate and bittersweet finish."
# }
# ]
```
### Add Beer review
``` bash
POST /beers/{id}/reviews

# Request sample
# {
#     "first_name":"Joe",
#     "last_name":"Tribiani",
#     "score":5,
#     "text":"This is good but this is not pizza!"
# }
```

## Getting started
These instructions will install MongoDB, load sample data into database, and install Beer-review API server.

### Using Docker

#### Prerequisites
* Docker
* Docker compose
* Update MongoDB username and password in `db_user.txt` and `db_password.txt`

#### Installing
``` bash
docker-compose up
```

### Using Kubernetes

#### Prerequisites
* Docker desktop or Minikube or Kubernetes cluster
* Update MongoDB username and password in `db_user.txt` and `db_password.txt`
* Create a kubernetes secret from cli
``` bash
kubectl create secret generic db-user-pass --from-file=./db_user.txt --from-file=./db_password.txt
```
#### Installing
``` bash
kubectl apply -f kubernetes/
```
