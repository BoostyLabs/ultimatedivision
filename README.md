## Deployment architecture

**Solution**.

To make deployment independent of cloud providers, use containerization through docker images. That could be runnable on any server, as result we could switch between providers whenever we want without changing the deployment process.

**File locations.**

All docker files should locate in **${projectname}/deploy** directory.
For each service, at the project, write a separate docker file.

**Naming.**

According to docker files naming convention, it should have name of service before dot (ex.: signer.Dockerfile, projectname.Dockerfile).

If the project has several docker-compose files, these files should also have naming according to docker files naming convention (docker-compose.test.yml, docker-compose.local.yml).

**Deployment.**

For deployment use GitHub actions that trigger commands from Makefile. It will build docker images (with commit hash and latest), and it will be pushed to our docker registry. Images from docker registry will use on deployment server in docker-compose file.

**Rollback to particular version.**

On deployment part,  create docker image with name that contains commit hash (docker-registry-address/service-name:commit-hash), as result we could rollback to particular version whenever we want.

**Installing.**

Golang is our backend language.

We are using version 1.17.4. You can download it from the official website [GoLang](https://go.dev/dl/), and install it according to the official [instructions](https://go.dev/doc/install.)

**Database.**

For our project we use a relational database PostgreSQL, version 12.11 which you can download by following the link from the official [website](https://www.postgresql.org/download/) or you can run your database in a Docker container.

**Docker.**

For isolated installation of databases and servers we need a Docker, version 20.10.16 or higher, you can download it at official [website](https://docs.docker.com/engine/install/)

```
docker run --name=db -e POSTGRES_PASSWORD='$YOUR_PASSWORD' -p $YOUR_PORTS -d --rm postgres
docker exec -it db createdb -U postgres ultimatedivision_test
```

**Run the main server.**

From the web/console directory at the root of the project use this commands to build front end part of project:
```
npm i --force && npm run build
```

There is a makefile to run the project, you can run it with the command in root of the project:
```
make run_local
```
After this you can open console on localhost:8088 and admin panel on localhost:8087

**Access to logs.**

For access to logs, we use [Dozzle](https://dozzle.dev/).
It's running as a separate service in docker-compose. To create login & password - pass as environment variables to docker-compose and provide credentials to QA and Devs. So that they have easy and fast access to logs.`

**Metrics & graphs.**

To collect standards (like CPU, Memory usage) or custom metrics we use [Prometheus](https://prometheus.io/docs/introduction/overview/).

To make graphs we use [Grafana](https://grafana.com/docs/grafana/latest/introduction/) which uses metrics passed by Prometheus.


**Metric examples.**
>metrics/metrics.go
```
type Metric struct {
   handler  http.Handler
   newUsers Counter
}
   
// NewUsersInc increment Counter newUsers.
func (metric *Metric) NewUsersInc() {
   metric.newUsers.Inc()
}

// NewMetric is a constructor for a Metric.
func NewMetric() *Metric {
    newUsers := prometheus.NewCounter(prometheus.CounterOpts{
        Name: "number_registrations",
        Help: "The total number of successful registrations.",
    })
    
    // Create a custom registry.
   registry := prometheus.NewRegistry()

   // Register using our custom registry gauge.
   registry.MustRegister(newUsers)
   return &Metric{
        // Expose metrics.
        handler:  promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}),
        newUsers: newUsers, 
    }
}
```
>console/consoleserver/controllers/auth.go
```
// Register creates a new user account.
func (auth *Auth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var err error
	var request users.CreateUserFields

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		auth.serveError(w, http.StatusBadRequest, AuthError.Wrap(err))
		return
	}

	if !request.IsValid() {
		auth.serveError(w, http.StatusBadRequest, AuthError.New("did not fill in all the fields"))
		return
	}

	err = auth.userAuth.Register(ctx, request.Email, request.Password, request.NickName, request.FirstName, request.LastName, request.Wallet)
	if err != nil {
		switch {
		case userauth.ErrAddressAlreadyInUse.Has(err):
			auth.serveError(w, http.StatusBadRequest, userauth.ErrAddressAlreadyInUse.Wrap(err))
			return
		default:
			auth.log.Error("Unable to register new user", AuthError.Wrap(err))
			auth.serveError(w, http.StatusInternalServerError, AuthError.Wrap(err))
			return
		}
	}
	
	auth.metric.LoginsInc()
}
```