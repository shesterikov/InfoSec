'use strict';

const escapeHtml = require('../utils/escape-html');

const layout = (title, content) => `<!doctype html><html lang="ru"><head><meta charset="utf-8"><title>${escapeHtml(title)}</title><style>body{font:16px system-ui;max-width:800px;margin:36px auto;padding:0 16px}a{margin-right:14px}input,textarea{padding:8px}button{padding:9px;margin:8px 0}.note{background:#eee;padding:12px}</style></head><body><nav><a href="/">Главная</a><a href="/search?q=hello">Reflected</a><a href="/comments">Stored</a><a href="/dom?name=Student">DOM</a><a href="/account">CSRF</a><a href="/frame">Clickjacking</a><a href="/api/private">API</a></nav><hr><h1>${escapeHtml(title)}</h1>${content}</body></html>`;

module.exports = layout;
