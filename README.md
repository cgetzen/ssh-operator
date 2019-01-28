ssh-operator allows you to ssh into a pod without `kubectl exec`.

# But... why?
["SSH is an anti-pattern."](https://twitter.com/bitfield/status/1062278396863041536)

Yes, ssh-operator is no-good, gross, and very very bad. But sometimes, no-good, gross, and very very bad is exactly what you need.

Ssh-operator was created to decouple "transitioning to kubernetes" and "taking away shell access", and could be useful when:
* You don't want to train your dev team with kubectl.
* You want to pair on a production system.
* You just want to see the world burn.


# Installation
Currently, you have to clone the repo and
```
kubectl apply -f deploy/ -n tmate
```

In the [future](https://github.com/cgetzen/ssh-operator/issues/2), you will be able to install this with helm.

# Directions
```
kubectl annotate pod $POD ssh.in="true"
```
After the ssh pod connects to $POD, you will be able to
```
kubectl get pod $POD -o=jsonpath="{.metadata.annotations.ssh}"
```
and use the connection string.


# What is tmate?
[Tmate](https://tmate.io/) is a fork of tmux, and can be used to share a screen.
In order to annotate a pod with an ssh address:
1. A new tmate pod is brought up
2. The tmate pod `kubectl exec`s to the original pod
3. The tmate pod annotates the original pod with the ssh address.

# Contributing
This is an easy kubernetes project to start contributing to! Take a look at the [issues](https://github.com/cgetzen/ssh-operator/issues) and [projects](https://github.com/cgetzen/ssh-operator/projects).
