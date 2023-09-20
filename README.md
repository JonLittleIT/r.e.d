# red-infra
red teaming infra and automations

Used for automated testing appsec tests and reporting.

Only set parameters for what each bounty program allows according to scope these are just examples of how this system can work for your company.

## Disclaimer: I am not responsible for you testing a companies security posture with out permission.


```mermaid
  graph TD;
      Sublister-->SubsFile;
      SubsFile-->httpx;
      httpxFile-->Nuclei;
      NucleiFileScanResults-->SlackWebhook;
```




```mermaid
gantt
    title A Gantt Diagram
    dateFormat YYYY-MM-DD
    section 1
        Attack surface tool auto task          :a1, 2023-09-10, 3d
        Armatage auto    :after a1, 20d
    section 2
        ML with AI tune of autos :2023-09-11, 12d
        Fine tune of results                 :30d

```

```
go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest
go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest
go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest
```
