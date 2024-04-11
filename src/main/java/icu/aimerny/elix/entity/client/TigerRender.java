package icu.aimerny.elix.entity.client;

import icu.aimerny.elix.Elix;
import icu.aimerny.elix.entity.custom.TigerEntity;
import net.minecraft.client.render.VertexConsumerProvider;
import net.minecraft.client.render.entity.EntityRendererFactory;
import net.minecraft.client.render.entity.MobEntityRenderer;
import net.minecraft.client.util.math.MatrixStack;
import net.minecraft.util.Identifier;

public class TigerRender extends MobEntityRenderer<TigerEntity, TigerModel<TigerEntity>> {
    public TigerRender(EntityRendererFactory.Context context) {
        super(context, new TigerModel<>(context.getPart(ModModelLayers.TIGER)),0.5F);
    }

    private static final Identifier TEXTURE = new Identifier(Elix.MOD_ID, "textures/entity/tiger.png");

    @Override
    public Identifier getTexture(TigerEntity entity) {
        return TEXTURE;
    }

    @Override
    public void render(TigerEntity mobEntity, float f, float g, MatrixStack matrixStack, VertexConsumerProvider vertexConsumerProvider, int i) {
        if(mobEntity.isBaby()) {
            matrixStack.scale(0.5F, 0.5F, 0.5F);
        }else{
            matrixStack.scale(1.5F, 1.5F, 1.5F);
        }
        super.render(mobEntity, f, g, matrixStack, vertexConsumerProvider, i);
    }
}
