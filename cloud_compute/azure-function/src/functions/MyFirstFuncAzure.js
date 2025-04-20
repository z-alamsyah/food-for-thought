const { app } = require('@azure/functions');
const { createClient } = require('@supabase/supabase-js');

app.http('MyFirstFuncAzure', {
    methods: ['GET', 'POST'],
    authLevel: 'function',
    handler: async (request, context) => {

        const supabase = createClient('https://lbsstqwalemvwlrnenjh.supabase.co', SUPA_KEY);
        let { data: todos, error } = await supabase
        .from('todos')
        .select()


        context.log(`get data todos:`, todos);

        return { body: JSON.stringify(todos) };
    }
});
