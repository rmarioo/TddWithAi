# TDD with AI pair programming experiment

This is a very simple experiment on doing tdd leveraging an AI as companion or as a pair programmer

To get a sense of this experiment is really reccommended to check the commits in [chronological order](https://github.com/rmarioo/TddWithAi/commits/master/) , so you can get how this particular TDD approach was done

Problem solved here is a subset of **Employee Report kata** https://codingdojo.org/kata/Employee-Report/ and i choose it on purpouse because is very simple

Here if you follow the commits in order i did TDD with ai in this way

- I manually created the failing tests
- I leverage AI to make test pass very quickly 
- I improved the code some time my self , some other time leveraging the AI

In [commits](https://github.com/rmarioo/TddWithAi/commits/master/) you can find also the prompts i used .

Why this approach?
Because i think that is critical to have tests that we can trust and can certify that our software is doing what we assume it should do.
Let the AI generate tests is a risk because at the moment we cannot really trust the result.

Also i found that AI is valuable as a companion that help us to quickly provide concrete examples of syntax 

For reference i used a dolphin-2.6-mistral-7b.Q8_0.gguf  LLM  that i was able to run in my local pc thanks to **LM Studio** and **codeGPT Intellij plugin**
