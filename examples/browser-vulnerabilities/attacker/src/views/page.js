'use strict';

const page=(title,body)=>`<!doctype html><html lang="ru"><meta charset="utf-8"><title>${title}</title><style>body{font:16px system-ui;max-width:750px;margin:40px auto}iframe{width:600px;height:240px}button{padding:10px}</style><h1>${title}</h1>${body}<p><a href="/">Назад</a></p>`;

module.exports = page;
