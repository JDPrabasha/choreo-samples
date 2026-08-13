'use strict';

const express = require('express');
const path = require('path');
const React = require('react');
const { renderToString } = require('react-dom/server');

const { createElement: e } = React;

const app = express();
const PORT = process.env.PORT || 8080;
const DEFAULT_DELAY_SECONDS = 90;

app.use(express.static(path.join(__dirname, 'public')));

function ReportPage() {
  return e(
    'html',
    null,
    e(
      'head',
      null,
      e('meta', { charSet: 'utf-8' }),
      e('title', null, 'Report Generator')
    ),
    e(
      'body',
      null,
      e(
        'div',
        { id: 'app' },
        e('h1', null, 'Report Generator'),
        e(
          'p',
          null,
          'This page is server-side rendered with React. Click the button to simulate ' +
            'generating a report that takes a long time to complete — used to test ' +
            `whether Choreo's endpoint timeout is enforced (default ${DEFAULT_DELAY_SECONDS}s).`
        ),
        e('button', { id: 'generate-btn' }, 'Generate Report'),
        e('p', { id: 'status' })
      ),
      e('script', { src: '/client.js' })
    )
  );
}

app.get('/', (req, res) => {
  const html = renderToString(e(ReportPage));
  res.type('html').send(`<!DOCTYPE html>${html}`);
});

// Simulates report generation that takes `seconds` (default 90) before responding,
// so a caller can observe whether the endpoint timeout cuts the request off first.
app.get('/api/generate-report', async (req, res) => {
  const parsed = parseInt(req.query.seconds, 10);
  const seconds = Number.isFinite(parsed) && parsed > 0 ? parsed : DEFAULT_DELAY_SECONDS;

  console.log(`Report generation requested, will take ${seconds}s`);
  const start = Date.now();
  await new Promise((resolve) => setTimeout(resolve, seconds * 1000));
  const elapsedSeconds = ((Date.now() - start) / 1000).toFixed(1);

  res.json({
    message: `Report generated after ${elapsedSeconds}s (requested ${seconds}s)`,
    generatedAt: new Date().toISOString(),
  });
});

app.listen(PORT, () => {
  console.log(`Server listening on port ${PORT}`);
});
