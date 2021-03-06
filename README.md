<h1 align="center">waioNEAT 🧬 🧠</h1>
<p align="center">
  <a title="Version" rel="nofollow"><img src="https://img.shields.io/static/v1?label=version&message=0.0.0&color=blue" alt="Version"></a>
 <p/>

# About
_Wow... Another Implementation Of NEAT_

There is plenty of NEAT implementations out there so here is mine.    
It's designed as a minimal and barebone library to add NEATs to your projects.  
Written in Go with no dependencies other than the go std.  

This is a pet project.

<p align="center">
  THIS IS MY NEAT. THERE ARE MANY LIKE IT BUT THIS ONE’S MINE....
<P/>

# Prerequisites
You probably need Go 1.18 or later. It may work for some older versions tho.

# Install
```
$ go get github.com/PauGalopaDev/waioNEAT@latest
```

# Usage  
Before anything else, initilize the module:
```go
  package main
  
  import(
    waio "github.com/PauGalopaDev/waioNEAT"
  )
  
  func main() int {
    waio.Init()
  }
```
*(Pending configuration by file)*

This module starts from the suposition that your program is going to have a set of parameters that will act as input and or outputs for your neural networks.

So we can start creating a genome:
```go
  ins := map[string]float64{
    "input1":1, // Name of the parameter and the chance of being present at the genome (from 0.0 to 1.0)
    "input2":0.5,
  }
  
  outs := map[string]float64{
    "output1":0.9,
    "output2":0.9,
  }

  genome := MakeGenome(ins, outs)
```

You can always get more creative with the names but its just a way to identify the purpose of each neuron.
Once we have the genome a network can be created:  

```go
  network := MakeNetwork(genome)
```

At this point, in order to Feed the network it can be done in various ways:

```go

  ins := map[string]float64{
    "input1":0.77, // Name of the parameter and its value
    "input2":0.55,
  }
  outs := network.FeedIn(ins)
  
  // Or Feed, which feeds the network with the current neuron values.
  // At this moment I don't know how to explain why this is useful but
  //  it has to do with pointers to Neuron.Value
  network.Feed()
  
  
```
  
Lastly, for genome crossover and mutation...
```go
  // As the function names says, it performs a crossover between genomes and returns its offspring.
  offspring := Crossover(genome1, genome2)
  
  // You can mutate the genome using:
  genome1.MutateAddNode()
  genome1.MutateAddSynapse()
  
  // A more realistic approach would be:
  mutationChance := 0.5
  for _, genome := range genomeList {
    if chance := rand.Float64(); chance <= mutationChance {
      genome.MutateAddSynapse()
    }
```

waioNEAT has the following types:  
  - Neuron
  - NeuronGene
  - SynapseGene
  - Genome
  - Network
  - ActivationFn

# Examples
_Pending_
