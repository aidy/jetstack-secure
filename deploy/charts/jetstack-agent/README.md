# jetstack-agent

Jetstack Secure Agent

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: v1.38.0](https://img.shields.io/badge/AppVersion-v1.38.0-informational?style=flat-square)

## Additional Information

The Jetstack secure agent helm chart installs the Kubernetes agent that connects to The TLS Protect For Kubernetes platform.
It will require a valid TLS Protect for Kubernetes organisation with a license to add the new cluster.
You should also choose a unique name for your cluster that it will appear under in the TLPK platform.

## Installing the Chart

### Obtaining credentials

First obtain your service account credential, this can be done through the UI or [jsctl](https://github.com/jetstack/jsctl/releases)

For example with `jsctl`:
```
jsctl set organization <MY_ORG>
jsctl auth login
jsctl auth clusters create-service-account <CLUSTER_NAME> | tee credentials.json
{
  "user_id": "SOME_AUTOGENERATED_USERID",
  "user_secret": "REDACTED"
}
```

### Deploying the chart

Once credentials are obtained, there are two ways to install the chart:

#### Method 1: create secret manually

```
# pre-create secret
kubectl create secret -n jetstack-secure "<SOME_SECRET_NAME>" --from-file=credentials.json
# Install refering to secret
helm upgrade --install --create-namespace -n jetstack-secure jetstack-agent \
  oci://eu.gcr.io/jetstack-secure-enterprise/charts/jetstack-agent \
  --set config.organisation="strange-jones"  --set config.cluster="<CLUSTER_NAME>" \
  --set authentication.secretName="<SOME_SECRET_NAME>"
```

#### Method 2: Pass secret to chart as a value, it creates the secret

*This is loading the secret obtained from create-service-account step [above](#obtaining-credentials)  
`export HELM_SECRET="$(cat credentials.json)"`*

```console
# Installing by passing in secret directly
helm upgrade --install --create-namespace -n jetstack-secure jetstack-agent \
  oci://eu.gcr.io/jetstack-secure-enterprise/charts/jetstack-agent \
  --set config.organisation="strange-jones" --set config.cluster="<CLUSTER_NAME>" \
  --set authentication.createSecret=true --set authentication.secretValue="$HELM_SECRET"
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| authentication.createSecret | bool | `false` |  |
| authentication.secretKey | string | `"credentials.json"` |  |
| authentication.secretName | string | `"agent-credentials"` |  |
| authentication.secretValue | string | `""` |  |
| authentication.type | string | `"file"` |  |
| config.cluster | string | `""` |  |
| config.dataGatherers.custom | list | `[]` |  |
| config.dataGatherers.default | bool | `true` |  |
| config.organisation | string | `""` |  |
| config.period | string | `"0h1m0s"` |  |
| config.server | string | `"https://platform.jetstack.io"` |  |
| fullnameOverride | string | `""` | Helm default setting, use this to shorten install name |
| image | object | `{"pullPolicy":"IfNotPresent","repository":"quay.io/jetstack/preflight","tag":"v0.1.38"}` | image settings |
| imagePullSecrets | list | `[]` | specify credentials if pulling from a customer registry |
| nameOverride | string | `""` | Helm default setting to override release name, leave blank |
| nodeSelector | object | `{}` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` | default replicas, do not scale up |
| resources.limits.cpu | string | `"500m"` |  |
| resources.limits.memory | string | `"500Mi"` |  |
| resources.requests.cpu | string | `"200m"` |  |
| resources.requests.memory | string | `"200Mi"` |  |
| securityContext.capabilities.drop[0] | string | `"ALL"` |  |
| securityContext.readOnlyRootFilesystem | bool | `true` |  |
| securityContext.runAsNonRoot | bool | `true` |  |
| securityContext.runAsUser | int | `1000` |  |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| serviceAccount.create | bool | `true` | Specifies whether a service account should be created @default true |
| serviceAccount.name | string | `""` |  |
| tolerations | list | `[]` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.10.0](https://github.com/norwoodj/helm-docs/releases/v1.10.0)
# jetstack-agent

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: v1.38.0](https://img.shields.io/badge/AppVersion-v1.38.0-informational?style=flat-square)

Jetstack Secure Agent

**Homepage:** <https://github.com/jetstack/jetstack-secure>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| JSCP and CRE Team | <tls-protect-for-kubernetes@jetstack.io> | <https://platform.jetstack.io/documentation> |

## Source Code

* <https://github.com/jetstack/jetstack-secure>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| authentication.createSecret | bool | `false` |  |
| authentication.secretKey | string | `"credentials.json"` |  |
| authentication.secretName | string | `"agent-credentials"` |  |
| authentication.secretValue | string | `""` |  |
| authentication.type | string | `"file"` |  |
| config.cluster | string | `""` |  |
| config.dataGatherers.custom | list | `[]` |  |
| config.dataGatherers.default | bool | `true` |  |
| config.organisation | string | `""` |  |
| config.period | string | `"0h1m0s"` |  |
| config.server | string | `"https://platform.jetstack.io"` |  |
| fullnameOverride | string | `""` | Helm default setting, use this to shorten install name |
| image | object | `{"pullPolicy":"IfNotPresent","repository":"quay.io/jetstack/preflight","tag":"v0.1.38"}` | image settings |
| imagePullSecrets | list | `[]` | specify credentials if pulling from a customer registry |
| nameOverride | string | `""` | Helm default setting to override release name, leave blank |
| nodeSelector | object | `{}` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` | default replicas, do not scale up |
| resources.limits.cpu | string | `"500m"` |  |
| resources.limits.memory | string | `"500Mi"` |  |
| resources.requests.cpu | string | `"200m"` |  |
| resources.requests.memory | string | `"200Mi"` |  |
| securityContext.capabilities.drop[0] | string | `"ALL"` |  |
| securityContext.readOnlyRootFilesystem | bool | `true` |  |
| securityContext.runAsNonRoot | bool | `true` |  |
| securityContext.runAsUser | int | `1000` |  |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| serviceAccount.create | bool | `true` | Specifies whether a service account should be created @default true |
| serviceAccount.name | string | `""` |  |
| tolerations | list | `[]` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.10.0](https://github.com/norwoodj/helm-docs/releases/v1.10.0)