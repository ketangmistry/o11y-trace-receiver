# o11y Trace Receiver
A containerized App which can be hosted on Google Cloud Platform, in particular Cloud Run. It will accept HTTP requests and publish the payload to a Pub/Sub topic. The GitHub action will create a container and push to Artefact Registry. There are some hooks to New Relic as well.
