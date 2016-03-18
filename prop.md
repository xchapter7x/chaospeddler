# Chaos-peddler 
## Service broker to employ chaos against your apps

** features **
* user binds app to service instance
* chaos-peddler now has app guid registered in its list.
* registered apps can be seen in the management portal and the chaos level can be adjusted.
* at the space level the instance is configured via a management portal for a certain level of chaos.
  * chaos levels:
    *   crazy - high frequency of AI killing, w/ high percentage of total AI count
    *   just-annoying high frequency of AI killing, w/ low percentage of total AI count
    *   mickey-mouse - lower frequency of AI killing, w/ lo percentage of total AI count 
    *   off - dont kill stuff

perhaps mvp we dont provide a management console we simply make the 4 chaos levels different plans.

---

** app healthcheck service**

- user binds app to service instance
- user is offered several plans 
  - healthcheck (pingdom style)
    - will this need a management console or will service params due?

---

** app attack service**

- user binds app to service instance
- user is offered several plans 
  - single 
    - a small plan offers only single source node attack
  - multi
    - a multi plan offers multiple concurrent source node attack
- user has management console where they can configure attack url, header, body, concurrency, time, etc. 
  - they can configure multiple attacks per service binding (will each be selected at random, each time or staggered?)
