package main

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"log"
)

func main() {
	d, err := diagram.New(diagram.Filename("aws"), diagram.Label("aws"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}

	ecs := aws.Compute.ElasticContainerService(diagram.NodeLabel("ECS"))
	lb := aws.Network.ElasticLoadBalancing(diagram.NodeLabel("LB"))

	d.Connect(lb, ecs)

	//dns := gcp.Network.Dns(diagram.NodeLabel("DNS"))
	//lb := gcp.Network.LoadBalancing(diagram.NodeLabel("NLB"))
	//cache := gcp.Database.Memorystore(diagram.NodeLabel("Cache"))
	//db := gcp.Database.Sql(diagram.NodeLabel("Database"))
	//
	//dc := diagram.NewGroup("GCP")
	//dc.NewGroup("services").
	//	Label("Service Layer").
	//	Add(
	//		gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 1")),
	//		gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 2")),
	//		gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 3")),
	//	).
	//	ConnectAllFrom(lb.ID(), diagram.Forward()).
	//	ConnectAllTo(cache.ID(), diagram.Forward())
	//
	//dc.NewGroup("data").Label("Data Layer").Add(cache, db).Connect(cache, db)

	//d.Add(ecs)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
