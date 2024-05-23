window.onload = function() {
    const defaultDefinitionUrl = window.location.host.includes('localhost') ? "http://localhost:8080/docs/users.yaml" : window.location.host.includes('dev.api.greeve.store') ? "https://dev.api.greeve.store/docs/users.yaml" : "https://api.greeve.store/docs/users.yaml";
    const ossServices = `petstore3.swagger.io=https://petstore3.swagger.io/api/v3/openapi.json,petstore31.swagger.io=https://petstore31.swagger.io/api/v31/openapi.json,petstore.swagger.io=https://petstore.swagger.io/v2/swagger.json,generator.swagger.io=https://generator.swagger.io/api/swagger.json,generator3.swagger.io=https://generator3.swagger.io/openapi.json,validator.swagger.io=https://validator.swagger.io/validator/openapi.json,oai.swagger.io=https://oai.swagger.io/api/openapi.json,converter.swagger.io=https://converter.swagger.io/api/openapi.json`;
    const ossServicesTuples = ossServices.split(',').map(ossService => ossService.split('='))
    const ossServiceMatch = ossServicesTuples.find(([host]) => window.location.host.includes(host))
    const definitionURL = ossServiceMatch ? ossServiceMatch[1] : defaultDefinitionUrl;

    
      //<editor-fold desc="Changeable Configuration Block">
      window.ui = SwaggerUIBundle({
        url: definitionURL,
        "dom_id": "#swagger-ui",
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout",
        queryConfigEnabled: true,
        validatorUrl: "https://validator.swagger.io/validator",
      })
      
      //</editor-fold>
};