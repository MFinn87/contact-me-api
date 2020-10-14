const nodemailer = require('nodemailer');
const Email = require('email-templates');
const path = require('path')
const root = path.join(__dirname, 'emails');

function sendEmail(name, phone, senderEmail, text) {
  const transporter = nodemailer.createTransport({
    service: process.env.emailService,
    auth: {
        user: process.env.emailUser,
        pass: process.env.emailAppKey
    }
    });

    const email = new Email({
        views: { root },
        message: {
          from: process.env.emailUser
        },
        // uncomment below to send emails in development/test env:
        send: true,
        transport: transporter,
        preview: false,
    })
      
      return email.send({
        template: 'template',
        message: {
          to: process.env.contactEmailRecipient
        },
        locals: {
        name: name,
        phone: phone,
        email: senderEmail,
        text: text,
        }
      })
}

module.exports = sendEmail