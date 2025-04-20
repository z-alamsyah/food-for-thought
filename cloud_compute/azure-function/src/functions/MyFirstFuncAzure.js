const { app } = require('@azure/functions');
const { createClient } = require('@supabase/supabase-js');

app.http('MyFirstFuncAzure', {
    methods: ['GET', 'POST'],
    authLevel: 'function',
    handler: async (request, context) => {

        const supabase = createClient('https://lbsstqwalemvwlrnenjh.supabase.co', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Imxic3N0cXdhbGVtdndscm5lbmpoIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTc0MTY3NDI2NiwiZXhwIjoyMDU3MjUwMjY2fQ.z6TIxCpEYv6QpVmIFMFnH92XwUYXcYlbd0SOphjtALA');
        let { data: todos, error } = await supabase
        .from('todos')
        .select()


        context.log(`get data todos:`, todos);

        return { body: JSON.stringify(todos) };
    }
});
