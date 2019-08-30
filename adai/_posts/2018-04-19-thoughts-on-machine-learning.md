---
title: "Thoughts on machine learning"
tags: 
  - machine learning
  - deep learning
categories:
  - guides
last_modified_at: 2018-04-19    
---

A diverse array of algorithms has been developed such as matrix factorization (e.g., for recommender systems), latent Dirichlet allocation (e.g., for topic modeling), Markov models (e.g., sequential information), association analysis (e.g., market basket analysis), convolutional neural networks (CNN) (e.g., image processing), recurrent neural networks (e.g., text processing), and etc., to cover the wide variety of machine learning problems.

Consider the CIFAR-10, CIFAR-100 (multiclass image classification), and SVHN (recognizing digits in street view house numbers) challenge. Although these challenges seem similar, significantly different convolutional neural network (CNN) structures have been developed to attain the top accuracy rank. Namely, fractional max-pooling (think of it as a form of pooling and data augmentation at the same time)[^1] for CIFAR-10, exponential linear units[^2] for CIFAR-100, and generalized pooling[^3] for SVHN.

Most machine learning solutions are problem specific and involve modifications to hyperparameters, number of layers, filter size, activation units, pooling structures, connection matrix between layers, etc. These require manual tweaking by hand, and some has heuristics which are hard to justify or to generalize.

It perplexes me that a single homogenous human brain, presumably with reprogrammable cortexes, is capable of all the above tasks, yet vastly different machine learning structures have been developed for each of them in practice.

I have long been captivated by the idea of an adaptive learning algorithm, consisting of replicated basic building blocks of intelligence, which is reprogrammable for different tasks and requires no manual tuning of architecture or hyperparameters.

In this webiste of mine, I hope to record my journey towards a general, if not universal, learning algorithm. I believe by mingling with and learning from the best minds -- you, my readers -- in the field of machine learning, will hasten our collective progress towards a general solution.

### References
[^1]: B. Graham, "Fractional Max-Pooling," CoRR, vol. abs/1412.6071, 2014.

[^2]: D.-A. Clevert, et al., "Fast and Accurate Deep Network Learning by Exponential Linear Units (ELUs)," CoRR, vol. abs/1511.07289, 2015.

[^3]: C.-Y. Lee, et al., "Generalizing Pooling Functions in Convolutional Neural Networks: Mixed, Gated, and Tree," CoRR, vol. abs/1509.08985, 2015.
