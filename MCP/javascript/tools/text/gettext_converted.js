/**
 * Page text by URL
 */

import fs from 'fs';
import path from 'path';
import os from 'os';

function getConfig() {
  const baseURL = process.env.API_BASE_URL;
  const bearerToken = process.env.API_BEARER_TOKEN;
  
  if (!baseURL || !bearerToken) {
    const configPath = path.join(os.homedir(), '.api', 'config.json');
    try {
      const configData = JSON.parse(fs.readFileSync(configPath, 'utf8'));
      return {
        baseURL: baseURL || configData.baseURL,
        bearerToken: bearerToken || configData.bearerToken
      };
    } catch (e) {
      throw new Error('Configuration not found. Please set API_BASE_URL and API_BEARER_TOKEN environment variables or create config file at ~/.api/config.json');
    }
  }
  
  return { baseURL, bearerToken };
}

export async function get_text(text_format, url, wait_for, proxy, country, custom_proxy, device, js_script, timeout, js_timeout, return_links, js, error_on_404, error_on_redirect, headers) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (text_format) params.append("text_format", text_format);
      if (url) params.append("url", url);
      if (wait_for) params.append("wait_for", wait_for);
      if (proxy) params.append("proxy", proxy);
      if (country) params.append("country", country);
      if (custom_proxy) params.append("custom_proxy", custom_proxy);
      if (device) params.append("device", device);
      if (js_script) params.append("js_script", js_script);
      if (timeout) params.append("timeout", timeout);
      if (js_timeout) params.append("js_timeout", js_timeout);
      if (return_links) params.append("return_links", return_links);
      if (js) params.append("js", js);
      if (error_on_404) params.append("error_on_404", error_on_404);
      if (error_on_redirect) params.append("error_on_redirect", error_on_redirect);
      if (headers) params.append("headers", headers);
    const queryString = params.toString();
    const finalUrl = queryString ? `${url}?${queryString}` : url;
    
    const url = `${config.baseURL}/api/unknown`;
    
    const response = await fetch(finalUrl, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${config.bearerToken}`,
        'Accept': 'application/json'
      }
    });
    
    if (!response.ok) {
      return `Failed to format JSON: ${response.status} ${response.statusText}`;
    }
    
    try {
      const result = await response.json();
      return JSON.stringify(result, null, 2);
    } catch (e) {
      return await response.text();
    }
    
  } catch (error) {
    return `Request failed: ${error.message}`;
  }
}

export function createGetTextTool() {
  return {
    definition: {
      name: 'get-text',
      description: 'Page text by URL',
      inputSchema: {
        type: 'object',
        properties: {
          text_format: {
            type: 'string',
            description: 'Format of the text response (plain by default). "plain" will return only the page body text. "json" and "xml" will return a json/xml with "title", "description" and "content" keys.'
          },
          url: {
            type: 'string',
            description: 'URL of the target page.'
          },
          wait_for: {
            type: 'string',
            description: 'CSS selector to wait for before returning the page content. Useful for pages with dynamic content loading. Overrides js_timeout.'
          },
          proxy: {
            type: 'string',
            description: 'Type of proxy, use residential proxies if your site restricts traffic from datacenters (datacenter by default). Note that residential proxy requests are more expensive than datacenter, see the pricing page for details.'
          },
          country: {
            type: 'string',
            description: 'Country of the proxy to use (US by default).'
          },
          custom_proxy: {
            type: 'string',
            description: 'Your own proxy URL to use instead of our built-in proxy pool in "http://user:password@host:port" format (<a target="_blank" href="https://webscraping.ai/proxies/smartproxy">Smartproxy</a> for example).'
          },
          device: {
            type: 'string',
            description: 'Type of device emulation.'
          },
          js_script: {
            type: 'string',
            description: 'Custom JavaScript code to execute on the target page.'
          },
          timeout: {
            type: 'number',
            description: 'Maximum web page retrieval time in ms. Increase it in case of timeout errors (10000 by default, maximum is 30000).'
          },
          js_timeout: {
            type: 'number',
            description: 'Maximum JavaScript rendering time in ms. Increase it in case if you see a loading indicator instead of data on the target page.'
          },
          return_links: {
            type: 'boolean',
            description: '[Works only with text_format=json] Return links from the page body text (false by default). Useful for building web crawlers.'
          },
          js: {
            type: 'boolean',
            description: 'Execute on-page JavaScript using a headless browser (true by default).'
          },
          error_on_404: {
            type: 'boolean',
            description: 'Return error on 404 HTTP status on the target page (false by default).'
          },
          error_on_redirect: {
            type: 'boolean',
            description: 'Return error on redirect on the target page (false by default).'
          },
          headers: {
            type: 'object',
            description: 'HTTP headers to pass to the target page. Can be specified either via a nested query parameter (...&headers[One]=value1&headers=[Another]=value2) or as a JSON encoded object (...&headers={"One": "value1", "Another": "value2"}).'
          }
        },
        required: ["url"]
      }
    },
    handler: get_text
  };
}