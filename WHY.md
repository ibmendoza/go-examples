via **https://www.finextra.com/blogposting/6836/7-reasons-why-software-development-is-so-hard** by Paul Smyth

**Introduction**

I’m often asked by lay people why we humans can undertake large construction or engineering projects with relative success and yet we struggle to deliver software projects without bugs and on time.

In an attempt to explain why this is the case I give below 7 reasons why software development is very difficult. This is not an attempt to condone shoddy software development practices. Rather, I’m trying to show that even with excellent development processes it is still difficult to do software development right.

**The software industry is young**

Humans have been building house, roads and bridges for thousands of years. We’ve no idea how many house or bridges collapsed in the early days as humans learned the correct techniques for building these structures.

One has only to look at the infamous Tahoma Narrows bridge collapse in 1940 to realise that, thousands of years after the first bridges were built, they still hadn’t perfected bridge building.

In comparison the software industry is only about 50 years old. We still have a long way to go before we have the body of experience behind us that the construction and manufacturing industries have.

Today the construction industry use mostly prefabricated materials and parts. Most of these are made by machine and have been tried and tested on many other projects.

The software industry on the other hand doesn’t have the range of pre-built components that other industries have. Software systems are fundamentally built by a process of discovery, invention, and creation of new components with the result that each new system is a custom project created from scratch. This leads us to our next point.

**Every line of code is a potential point of failure**

Because all new projects are custom built it follows that every line of code is unproven and therefore should be tested. However, in the real world, this is totally impractical.

Each line of code will have dozens, even thousands, of possible inputs, outputs, states or dependencies to deal with. It can impact, or be impacted by, other lines of code or by external factors. Even if it was possible to document every test case for a line of code you still couldn’t be sure that there wasn’t some unknown factor that could cause an error.

And testing a single line of code is only part of the challenge. No line of code exists on its own. It is part of the whole system and the whole needs to be tested to ensure that all parts of the application function correctly.

The sheer complexity of software means it is impossible to test every path so in the real world the best project teams will implement processes that are designed to increase the likelihood of the software being defect free. They will use techniques such as coding standards, unit testing, smoke testing, automated regression testing, design and code reviews etc. all of which should improve the quality of the software.

All of this testing comes at a cost. The question to be answered on every project is – how critical is this software and how much testing should we do to ensure the software is correct?

Too often the testing phase is rushed and the software goes out with an unacceptable level of defects. On the other hand, for most systems there are diminishing returns for extending the testing past a certain point. There comes a point with all software where the value of getting the software released is greater than the value gained by continuing to test for defects. This is why most commercial software gets released even though it is known to contain defects.

**Lack of user input**

For over 10 years the research company, The Standish Group, have surveyed companies about their IT projects. The No. 1 factor that caused software projects to become challenged was "Lack of User Input".

Reasons for this can include:

    The system is being promoted by the management and so the business users have no buy-in
    The users are too busy and have “more important” things to do
    Relations between the user community and the I.T. team are poor 

Without the involvement and input of a user representative the project is doomed to failure. This person should be a subject domain expert with the authority to make decisions and a commitment to the project timescales.

So assuming there is good user input then the challenge of translating requirements into a design begins. And this is no easy task as our next point shows.

**Users don't know what they want until they see it**

Even with good input from the users no amount of analysis of user requirements can take away an immutable fact that users only think that they know what they want. In truth, it’s not until they start seeing something, and using it, that they begin to really understand what they need. This is especially true when the software is being developed for a new idea or process that they haven’t used before.

Studies have shown that the average project experiences about a 25% change in requirements from the "requirements complete" stage to the first release. This is the famous “scope creep” problem that besets nearly all projects. It usually starts once the first designs begin to appear which cause the users to think more deeply about what they really want.

The challenge is do you a) ignore new requirements and carry on building according to the design documents and risk delivering a system that doesn’t do what the users need or b) take on changes as they arise with the result that the project expands and costs rise?

There is no simple answer to this dilemma despite the fact that various techniques, such as Agile development, have evolved to make it easier to adapt to changing requirements. Even seemingly small changes can have a big impact on a project. Each explicit requirement can lead to many more implicit requirements (by a factor of up to 50) further complicating the software.

Changing requirements during the development phase is one of the great challenges facing all software developers. It can be done but don’t ever think it’s easy. And please don’t think new requirements can be accommodated without affecting the timeline or the budget unless there is a corresponding removal of requirements.

**There are no barriers to entry to become a programmer**

There is one argument that states that software development is so hard because programming is so easy. In other words it is relatively easy to learn how to write code but there is a huge gap between that and delivering great software.

One could possibly equate it to learning a new language. Yes you may pick up the grammar and acquire a reasonable vocabulary but that’s a whole different ball game to having a fluent two-way discussion with some native speakers.

Various studies have shown that the productivity ratio between different grades of developer can be as high as 28:1. With that in mind it you surely would only want to hire the best developers. Unfortunately this is not easy as great developers are a very rare commodity.

There is no barrier to entry into the programming world and thus it is awash with many poor programmers who adversely affect projects. In addition, even potentially good young developers will still make mistakes that a more experienced developer will have learned to avoid.

It really is worth paying more for a top-class experienced developer. They will do things quicker, and better and with less code. Your project will be delivered quicker and will have fewer defects. They will save you money now and they will also save you money through the life of the system in support and maintenance costs.

**All software is affected by external factors**

Physical structures obey physical laws e.g. they are affected by gravity, mass, atmosphere etc. Through thousands of years of learning much is known about the physical world and can therefore be modelled and predicted.

Software is “mindware” and therefore doesn’t obey physical laws but it usually must conform to external constraints such as hardware, integration with other software, government regulations, legacy data formats, performance criteria, scalability etc.

Understanding and catering for all of these external factors is a near impossible task. Even a seemingly simple requirement, such as supporting multiple browsers, exponentially increases the difficulty of both building and testing software. If you then add in a requirement to support multiple versions of each browser then you are again exponentially increasing the complexity and the difficulty.

**Estimating is an art not a science**

Given the premise that all new projects are custom built, that there are no pre-built components, that the project will suffer from scope creep and that the level of skill across the development team is usually varied then it is no wonder that estimating the duration of the project can never be a scientific exercise. There are too many unknowns. If you’ve never built the same software before with the same team then how can you know how long it will take?

Of course experience guides you in your estimating and the more experience you have the more likely you will be to anticipate the unknowns. Too many projects run over because overly optimistic estimates are set by inexperienced people who expect everything to flow smoothly and who make no allowance for the unknowns.

Over the last 10 years the Agile Development method has emerged as a means of addressing these estimating issues. Whilst this is generally a more reliable way to control projects timescales, and therefore the cost, it is not suitable for all projects or even on some parts of projects.

Where projects involve complex external interfaces or new technology is being used then estimates become even harder to get right. These are risks that are often hard to quantify up-front and are usually only uncovered as the work gets done.

**Summary**

A software application is like an iceberg – 90% of it is not visible. The major complexity in the application lies below the waterline and is invisible to the user. The next time you are wondering why software projects appear so difficult to get right you might perhaps spare a thought for the developers. They are generally working hard to deliver on time against a tidal wave of challenges and complexity.
