# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security
Name: e_sti_ca_crl_distribution\
Source: ATIS-1000080v4\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 18 Oct 21 00:00 UTC\
Description: STI intermediate certificates shall contain a CRL Distribution Points extension containing a single DistributionPoint entry

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|

no warning, or error, or not effective date level issues were found

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | CN=Martini Security SHAKEN G1, O=Martini Security\\, LLC, L=Seattle, ST=WA, C=US | [view](../0fe086f321e93ca9ae08a19a89bf9049b7625fcf/README.md) | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| error | CN=Martini Security SHAKEN G2, O=Martini Security\\, LLC, L=Seattle, ST=WA, C=US | [view](../6af848efe4c680e44114ee32350826704eca8135/README.md) | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |


Generated: 26/10/2022 at 21:01:13